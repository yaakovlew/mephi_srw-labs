<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary flex justify-between items-center">
      Вычислите функции принадлежности для левых частей приведенных правил
      <q-toggle
        v-model="addInfoToggle"
        color="yellow"
        label="Доп. информация"
        icon="lightbulb"
      />
    </div>
    <div v-if="addInfoToggle" class="text-h6 text-primary">
      <div class="text-main">Правила:</div>
      <ul class="q-py-none q-my-none">
        <li v-for="rule in variant?.rule" :key="rule.name">
          {{ rule.name }}
        </li>
      </ul>
    </div>
    <q-table
      flat
      bordered
      title="Ответ"
      :rows="rows"
      :columns="columns"
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
                mask="#.##"
              />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-btn
      v-if="!answer"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      flat
      bordered
      title="Правильный ответ"
      :rows="rowsAnswer"
      :columns="columnsAnswer"
      row-key="name"
      binary-state-sort
      hide-bottom
      :separator="'cell'"
      v-else
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
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div v-if="answer">
      Количество полученных баллов за этап: {{ answer.percentage }} из
      {{ answer.max_mark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Ref, computed, onMounted, onUnmounted, ref, watch } from 'vue';
import { useLab3bStore } from '../../../stores/lab3b';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLab3bStore();

const isDone = computed(() => labStore.info?.step !== 1);

let columns = ref([]);

const addInfoToggle = ref(false);

const rows = ref([]);

onMounted(() => {
  if (!props.variant) return;
  columns.value = props.variant?.alternative.map((alternative, i) => {
    return {
      name: i,
      required: true,
      label: `A${i + 1}`,
      style: 'width: 300px',
    };
  });

  columns.value.unshift({
    name: 'name',
    required: true,
    label: 'Название',
  });

  const userAnswers = localStorage.getItem('user-answers-2');
  if (userAnswers) {
    rows.value = JSON.parse(userAnswers);
  } else {
    props.variant?.rule.forEach((criteria, i) => {
      const row = {};

      columns.value.forEach((column) => {
        const name = column.name;
        row.name = `d${i + 1}`;
        row[name] = 0;
      });

      rows.value.push(row);
    });
  }

  if (labStore.info?.step === 1) {
    localStorage.removeItem('answer-2');
    answer.value = null;
  }
});

watch(
  rows,
  () => {
    localStorage.setItem('user-answers-2', JSON.stringify(rows.value));
  },
  {
    deep: true,
  }
);

const isSend = ref(false);
const answer: Ref<Lab.AlternativeSetAnswer | null> = ref(
  localStorage.getItem('answer-2')
    ? JSON.parse(localStorage.getItem('answer-2')!)
    : null
);

const sendResult = async () => {
  const res = rows.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  const resAnswer = await labStore.sendRuleNumber({ matrix: res });
  if (resAnswer.data) {
    answer.value = resAnswer.data;
    localStorage.setItem('answer-2', JSON.stringify(answer.value));
    labStore.getInfo();
  }
};

const columnsAnswer = ref([]);
const rowsAnswer = ref([]);

watch(answer, () => {
  if (answer.value) {
    columnsAnswer.value.push({
      name: 'name',
      required: true,
      label: 'Название',
    });

    const answerColumnsLenght = answer.value?.result[0].length;
    for (let i = 0; i < answerColumnsLenght; i++) {
      columnsAnswer.value.push({
        name: i,
        required: true,
        label: `A${i + 1}`,
        style: 'width: 300px',
      });
    }

    rowsAnswer.value = answer.value?.result.map((row, i) => {
      const resRow = [];
      for (let a in row) {
        if (a !== 'name') {
          resRow.push(row[a] === -1 ? '' : row[a]);
        }
      }
      resRow.name = { data: `d${i + 1}` };
      return resRow;
    });
  }
});

onMounted(() => {
  if (!answer.value) return;
  columnsAnswer.value.push({
    name: 'name',
    required: true,
    label: 'Название',
  });

  const answerColumnsLenght = answer.value?.result[0].length;
  for (let i = 0; i < answerColumnsLenght; i++) {
    columnsAnswer.value.push({
      name: i,
      required: true,
      label: `A${i + 1}`,
      style: 'width: 300px',
    });
  }

  rowsAnswer.value = answer.value?.result.map((row, i) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === -1 ? '' : row[a]);
      }
    }
    resRow.name = { data: `d${i + 1}` };
    return resRow;
  });
});
</script>

<style lang="scss" scoped></style>
