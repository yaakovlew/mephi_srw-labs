<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary flex justify-between items-center">
      <div>
        Определите степень принадлежности альтернатив к классам по критерию
      </div>
      <q-toggle
        v-model="addInfoToggle"
        color="yellow"
        label="Доп. информация"
        icon="lightbulb"
      />
    </div>
    <div v-if="addInfoToggle" class="text-h6 text-primary flex justify-between">
      <div v-for="(info, i) in addInfo" :key="i">
        <div>
          {{ info.name }}
        </div>
        <ul>
          <li v-for="(item, j) in info.funcMark" :key="j">
            {{ item.name }}:
            {{ item.func }}
          </li>
        </ul>
      </div>
    </div>
    <template v-for="(table, index) in columns" :key="index">
      <div class="banner-title" v-if="index < currentStep">
        {{ `C${index + 1}` }}
      </div>
      <div class="flex g-m column" v-if="index < currentStep">
        <q-table
          flat
          bordered
          title="Ответ"
          :rows="table.rows"
          :columns="table.cols"
          row-key="name"
          binary-state-sort
          hide-bottom
          :separator="'cell'"
        >
          <template v-slot:body="props">
            <q-tr :props="props">
              <q-td v-for="col in props.cols" :key="col.name" :props="props">
                {{ props.row[col.name] === -1 ? '' : props.row[col.name] }}
              </q-td>
            </q-tr>
          </template>
        </q-table>
        <q-table
          flat
          bordered
          title="Правильный ответ"
          :rows="columnsAnswer[index].rows"
          :columns="columnsAnswer[index].cols"
          row-key="name"
          binary-state-sort
          hide-bottom
          :separator="'cell'"
        >
          <template v-slot:body="props">
            <q-tr :props="props">
              <q-td
                v-for="col in props.cols"
                :key="col.name"
                :props="props"
                :class="{ 'error-field': props.row[col.name].flag === false }"
              >
                {{ props.row[col.name] === -1 ? '' : props.row[col.name].data }}
                <q-popup-edit
                  v-if="col.name !== 'name' && isDone === false"
                  v-model="props.row[col.name]"
                  buttons
                  v-slot="scope"
                >
                  <q-input
                    v-model.number="scope.value"
                    @blur="scope.value = roundToNDigits(scope.value)"
                    type="number"
                    dense
                    autofocus
                  />
                </q-popup-edit>
              </q-td>
            </q-tr>
          </template>
        </q-table>
        <div v-if="answer[index]">
          Количество полученных баллов за шаг: {{ answer[index].percentage }} из
          {{ answer[index].max_mark }}
        </div>
      </div>
    </template>
    <div class="banner-title" v-if="columns[currentStep]?.rows">
      {{ `C${currentStep + 1}` }}
    </div>
    <q-table
      v-if="columns[currentStep]?.rows"
      flat
      bordered
      title="Ответ"
      :rows="columns[currentStep]?.rows"
      :columns="columns[currentStep]?.cols"
      row-key="name"
      binary-state-sort
      hide-bottom
      :separator="'cell'"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-for="col in props.cols" :key="col.name" :props="props">
            {{ props.row[col.name] === -1 ? '' : props.row[col.name] }}
            <q-popup-edit
              v-if="col.name !== 'name' && isDone === false"
              v-model="props.row[col.name]"
              buttons
              v-slot="scope"
            >
              <q-input
                v-model.number="scope.value"
                @blur="scope.value = roundToNDigits(scope.value)"
                type="number"
                dense
                autofocus
              />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-btn
      v-if="!isDone && currentStep < columns.length"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <div v-if="answer" style="font-size: 18px">
      Общее количество полученных баллов за шаги: {{ totalSum.sum }} из
      {{ totalSum.maxCount }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useLab3bStore } from '../../../stores/lab3b';
import { Lab3b } from 'src/models/lab/lab3b';
import { lab3bService } from '../../../services/lab3b';
import { data } from 'autoprefixer';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
  next_matrices_headers: string[][];
}>();

const labStore = useLab3bStore();

const isDone = computed(() => labStore.info?.step !== 0);

const addInfoToggle = ref(false);

const addInfo = computed(() =>
  props.variant.criteria.map((cr) => ({
    name: cr.definition.slice(0, 2),
    funcMark: cr.func_mark,
  }))
);

let columns = ref([]);

const columnsToCreateTable = ref('');
const columnsCount = ref(0);

const currentStep = ref(0);

const createTable = (colsTemp: string[]) => {
  const cols = colsTemp
    .map((col) => {
      let finalCol = col;

      return finalCol;
    })
    .filter((col) => col);

  columnsToCreateTable.value = '';
  const tableCols: any[] = [
    {
      name: 'name',
      label: 'Название',
      field: 'name',
    },
  ];
  const tableRows = [];

  cols.forEach((col, i) =>
    tableCols.push({ name: col, label: col, style: 'width: 400px' })
  );
  for (let i = 0; i < props.variant?.rule.length; i++) {
    const row = {
      name: `A${i + 1}`,
    };
    cols.forEach((col) => (row[col] = 0));
    tableRows.push(row);
  }

  columns.value.push({ cols: tableCols, rows: tableRows, done: false });
};

const deleteTable = (index: number) => {
  columns.value.splice(index, 1);
};

watch(
  columns,
  () => {
    localStorage.setItem('user-answers-1', JSON.stringify(columns.value));
  },
  {
    deep: true,
  }
);

const rows = ref([]);

onMounted(() => {
  if (!props.variant) return;

  const userAnswers = localStorage.getItem('user-answers-1');
  if (userAnswers) {
    columns.value = JSON.parse(userAnswers);
  } else {
    props.next_matrices_headers?.map((row) => createTable(row));
  }

  if (localStorage.getItem('step-1')) {
    currentStep.value = parseInt(localStorage.getItem('step-1')) ?? 0;
  }

  if (currentStep.value === 0) {
    localStorage.removeItem('answer-1');
    answer.value = [];
  }
});

watch(
  rows,
  () => {
    localStorage.setItem('user-answers-1', JSON.stringify(rows.value));
  },
  {
    deep: true,
  }
);

watch(
  currentStep,
  () => {
    localStorage.setItem('step-1', JSON.stringify(currentStep.value));
  },
  {
    deep: true,
  }
);

const totalSum = computed(() => {
  let sum = 0;
  let maxCount = 0;
  answer.value?.forEach((item) => {
    if (item) {
      sum += item.percentage;
      maxCount += item.max_mark;
    }
  });

  localStorage.setItem(
    'answer-1-sum',
    JSON.stringify({
      percentage: sum,
      max_mark: maxCount,
    })
  );

  return {
    sum,
    maxCount,
  };
});

const isSend = ref(false);
const answer: Ref<Lab3b.AlternativeSetAnswer[] | null> = ref(
  localStorage.getItem('answer-1')
    ? JSON.parse(localStorage.getItem('answer-1')!)
    : []
);

const sendResult = async () => {
  if (currentStep.value === columns.value.length) {
    return;
  }
  const rows = columns.value[currentStep.value].rows.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });
  const colsWithoutName = columns.value[currentStep.value]?.cols.filter(
    (col) => col.name !== 'name'
  );
  const result = {};
  colsWithoutName.forEach((c, i) => {
    result[c.name] = [];
    rows.forEach((r) => {
      result[c.name].push(Number(r[i]));
    });
  });

  const resAnswer = await labStore.sendRuleValue({
    matrices: result,
    step: currentStep.value + 1,
  });
  if (resAnswer.data) {
    currentStep.value = currentStep.value + 1;
    answer.value?.push(resAnswer.data);
    localStorage.setItem('answer-1', JSON.stringify(answer.value));
    if (currentStep.value === columns.value.length) {
      await lab3bService.increment0Step();
    }
    await labStore.getInfo();
  }
};

const columnsAnswer = ref([]);

watch(
  answer,
  () => {
    if (answer.value) {
      console.log(answer.value);
      columnsAnswer.value = answer.value?.map((arr) => {
        const colsName = Object.keys(arr.result);

        const cols = colsName.map((col) => {
          return {
            name: col,
            label: col,
            style: 'width: 400px',
          };
        });

        cols.unshift({
          name: 'name',
          label: 'Название',
        });
        const rows = [];
        rows.length = Object.values(arr.result)[0].length;
        rows.fill();
        Object.entries(arr.result).map(({ 0: key, 1: value }) => {
          value.forEach((cell, i) => {
            rows[i] = { ...rows[i], [key]: cell, name: { data: `A${i + 1}` } };
          });
        });

        return { cols: cols, rows: rows };
      });
    }
  },
  { deep: true }
);

onMounted(() => {
  if (!answer.value) return;
  columnsAnswer.value = answer.value?.map((arr) => {
    const colsName = Object.keys(arr.result);

    const cols = colsName.map((col) => {
      return {
        name: col,
        label: col,
        style: 'width: 400px',
      };
    });

    cols.unshift({
      name: 'name',
      label: 'Название',
    });
    const rows = [];
    rows.length = Object.values(arr.result)[0].length;
    rows.fill({});
    Object.entries(arr.result).map(({ 0: key, 1: value }) => {
      value.forEach((cell, i) => {
        rows[i] = { ...rows[i], [key]: cell, name: { data: `A${i + 1}` } };
      });
    });

    return { cols: cols, rows: rows };
  });
});
</script>

<style lang="scss" scoped></style>
