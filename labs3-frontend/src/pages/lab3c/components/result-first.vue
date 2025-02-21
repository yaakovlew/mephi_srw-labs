<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary flex justify-between items-center">
      <div>
        Определите оценки альтернатив по критериям с помощью заданных функций
        принадлежности
      </div>
      <q-toggle
        v-model="addInfoToggle"
        color="yellow"
        label="Доп. информация"
        icon="lightbulb"
      />
    </div>
    <div
      v-if="addInfoToggle"
      class="text-h6 text-primary flex column g-m justify-between"
    >
      <table>
        <tbody>
          <tr>
            <th class="row-info" scope="row">очень низкий</th>
            <td class="row-info">ОН</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">низкий</th>
            <td class="row-info">Н</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">средний</th>
            <td class="row-info">С</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">высокий</th>
            <td class="row-info">В</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">очень высокий</th>
            <td class="row-info">ОВ</td>
          </tr>
        </tbody>
      </table>
      <div class="text-h5 text-primary" style="font-weight: 700">
        Использовать данные сокращенные варианты
      </div>
    </div>
    <q-table
      flat
      bordered
      title="Ответ"
      :rows="rows"
      :columns="columnsFirst"
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
              v-if="
                col.name !== 'name' && isDone === false && currentStep === 0
              "
              v-model="props.row[col.name]"
              buttons
              v-slot="scope"
            >
              <q-input v-model="scope.value" />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-btn
      v-if="currentStep === 0"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="currentStep > 0"
      flat
      bordered
      title="Правильный ответ"
      :rows="rowsAnswer"
      :columns="columnsFirst"
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
            {{
              props.row[col.name] === -1
                ? ''
                : props.row[col.name].data
                ? String(props.row[col.name].data).toLocaleUpperCase()
                : props.row[col.name]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-table
      v-if="currentStep !== 0"
      flat
      bordered
      title="Ответ"
      :rows="rows"
      :columns="columnsSecond"
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
              v-if="
                col.name !== 'name' && isDone === false && currentStep === 1
              "
              v-model="props.row[col.name]"
              buttons
              v-slot="scope"
            >
              <q-input v-model="scope.value" />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-btn
      v-if="currentStep === 1"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="currentStep > 1"
      flat
      bordered
      title="Правильный ответ"
      :rows="rowsAnswer"
      :columns="columnsSecond"
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
            {{
              props.row[col.name] === -1
                ? ''
                : props.row[col.name].data
                ? String(props.row[col.name].data).toLocaleUpperCase()
                : props.row[col.name]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-table
      v-if="currentStep > 1"
      flat
      bordered
      title="Ответ"
      :rows="rows"
      :columns="columnsThird"
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
              v-if="
                col.name !== 'name' && isDone === false && currentStep === 2
              "
              v-model="props.row[col.name]"
              buttons
              v-slot="scope"
            >
              <q-input v-model="scope.value" />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-btn
      v-if="currentStep === 2"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="currentStep === 3"
      flat
      bordered
      title="Правильный ответ"
      :rows="rowsAnswer"
      :columns="columnsThird"
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
            {{
              props.row[col.name] === -1
                ? ''
                : props.row[col.name].data
                ? String(props.row[col.name].data).toLocaleUpperCase()
                : props.row[col.name]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div v-if="answer?.length">
      Общее количество полученных баллов за шаги: {{ totalSum.sum }} из
      {{ totalSum.maxCount }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Ref, computed, onMounted, onUnmounted, ref, watch } from 'vue';
import { useLab3cStore } from 'src/stores/lab3c';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const addInfoToggle = ref(false);

const labStore = useLab3cStore();

const isDone = computed(() => labStore.info?.step !== 0);

const currentStep = ref(0);

let columns = ref([]);
let columnsToShow = ref([]);

let columnsFirst = ref([]);
let columnsSecond = ref([]);
let columnsThird = ref([]);

const rows = ref([]);

const variants = ['В', 'С', 'Н', 'ОН'];

onMounted(() => {
  if (!props.variant) return;
  columns.value = props.variant?.alternative.map((alternative, i) => {
    return {
      name: i,
      required: true,
      label: `a${i + 1}`,
      align: 'left',
    };
  });

  columns.value.unshift({
    name: 'name',
    required: true,
    label: 'Название',
    style: 'width: 50%',
    align: 'left',
  });

  const userAnswers = localStorage.getItem('user-answers-1');
  if (userAnswers) {
    rows.value = JSON.parse(userAnswers);
  } else {
    props.variant?.criteria.forEach((criteria, i) => {
      const row = {};

      columns.value.forEach((column) => {
        const name = column.name;
        row.name = `С${i + 1}`;
        row[name] = variants[3];
      });

      rows.value.push(row);
    });
  }

  const localAnswers = localStorage.getItem('rowsAnswer-1');
  if (localAnswers) {
    rowsAnswer.value = JSON.parse(localAnswers);
  } else {
    rowsAnswer.value = JSON.parse(JSON.stringify(rows.value));
  }

  if (localStorage.getItem('step-1')) {
    currentStep.value = JSON.parse(localStorage.getItem('step-1')!);
  } else {
    currentStep.value = 0;
  }

  columnsToShow.value = columns.value.filter(
    (col, i) => i <= currentStep.value + 1 || col.name === 'name'
  );

  columnsFirst.value = columns.value.filter(
    (col, i) => i === 1 || col.name === 'name'
  );
  columnsSecond.value = columns.value.filter(
    (col, i) => i === 2 || col.name === 'name'
  );
  columnsThird.value = columns.value.filter(
    (col, i) => i === 3 || col.name === 'name'
  );

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
    columnsToShow.value = columns.value.filter(
      (col, i) => i <= currentStep.value + 1 || col.name === 'name'
    );
  },
  {
    deep: true,
  }
);

const isSend = ref(false);
const answer: Ref<Lab.AlternativeSetAnswer[] | null> = ref(
  localStorage.getItem('answer-1')
    ? JSON.parse(localStorage.getItem('answer-1')!)
    : null
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

const sendResult = async () => {
  const res = [];
  rows.value.map((row) => {
    for (let a in row) {
      if (a === String(currentStep.value)) {
        res.push(String(row[a]).toLowerCase());
      }
    }
  });

  const resAnswer = await labStore.sendAlternativeMatrix({
    matrix: res,
    step: currentStep.value + 1,
  });
  if (resAnswer.data) {
    answer.value?.push(resAnswer.data);
    currentStep.value = currentStep.value + 1;
    localStorage.setItem('answer-1', JSON.stringify(answer.value));
    if (currentStep.value + 1 === columns.value.length) {
      await labStore.increment0Step();
      labStore.normalStep = 2;
    }
    labStore.getInfo();
  }
};

const columnsAnswer = ref([]);
const rowsAnswer = ref([]);

watch(
  rowsAnswer,
  () => {
    localStorage.setItem('rowsAnswer-1', JSON.stringify(rowsAnswer.value));
  },
  {
    deep: true,
  }
);

watch(
  answer,
  () => {
    if (answer.value?.length) {
      columnsAnswer.value = columns.value.filter(
        (col, i) => i <= currentStep.value || col.name === 'name'
      );
      answer.value[currentStep.value - 1]?.result.map((row, i) => {
        rowsAnswer.value[i][currentStep.value - 1] = row;
      });
    }
  },
  {
    deep: true,
  }
);

onMounted(() => {
  if (!answer.value) return;
  columnsAnswer.value = columns.value.filter(
    (col, i) => i <= currentStep.value || col.name === 'name'
  );
  answer.value[currentStep.value - 1]?.result.map((row, i) => {
    rowsAnswer.value[i][currentStep.value - 1] = row;
  });
});
</script>

<style lang="scss" scoped>
.row-info {
  width: 50%;
}

.row-info {
  border: 1px solid var(--q-primary);
  padding: 8px;
  text-align: left;
  text-transform: capitalize;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
  border: 2px solid var(--q-primary);
}
</style>
