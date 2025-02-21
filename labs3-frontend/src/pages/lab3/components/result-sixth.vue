<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary">
      Постройте множество недоминирующих альтернатив Q2 (с учетом весов)
    </div>
    <q-table
      flat
      bordered
      title="Ответы"
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
              v-if="col.name !== 'name' && !isDone"
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
            {{ props.row[col.name].data ?? props.row[col.name] }}
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
import { computed, onMounted, ref, watch } from 'vue';
import { useLabStore } from '../../../stores/lab';
import { Ref } from 'vue';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLabStore();

const isDone = computed(() => labStore.info?.step !== 5);

const columns = props.variant?.alternative.map((alternative, i) => {
  return {
    name: i,
    required: true,
    label: `A${i + 1}`,
    style: 'width: 300px;',
  };
});

const rows = ref([]);

onMounted(() => {
  const row = {};

  columns.forEach((column) => {
    const name = column.name;
    row[name] = 0;
  });

  const userAnswers = localStorage.getItem('user-answers-6');
  if (userAnswers) {
    rows.value = JSON.parse(userAnswers);
  } else {
    rows.value.push(row);
  }

  if (labStore.info?.step === 5) {
    localStorage.removeItem('answer-6');
    answer.value = null;
  }
});

watch(
  rows,
  () => {
    localStorage.setItem('user-answers-6', JSON.stringify(rows.value));
  },
  {
    deep: true,
  }
);

const answer: Ref<Lab.NonDominatedAnswer | null> = ref(
  localStorage.getItem('answer-6')
    ? JSON.parse(localStorage.getItem('answer-6') || '')
    : null
);

const sendResult = async () => {
  const answ = rows.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  const res = await labStore.sendSecondNonDominated({ set: answ[0] });
  if (res.data) {
    answer.value = res.data;
    localStorage.setItem('answer-6', JSON.stringify(answer.value));
    labStore.getInfo();
  }
};

const columnsAnswer = ref([]);
const rowsAnswer = ref([]);

watch(answer, () => {
  if (answer.value) {
    const answerColumnsLenght = answer.value?.result.length;
    for (let i = 0; i < answerColumnsLenght; i++) {
      columnsAnswer.value.push({
        name: i,
        required: true,
        label: `A${i + 1}`,
        style: 'width: 300px;',
      });
    }

    const resRow = {};

    answer.value?.result.map((row, i) => {
      resRow[i] = row;
    });

    rowsAnswer.value.push(resRow);
  }
});

onMounted(() => {
  if (!answer.value) return;
  const answerColumnsLenght = answer.value?.result.length;
  for (let i = 0; i < answerColumnsLenght; i++) {
    columnsAnswer.value.push({
      name: i,
      required: true,
      label: `A${i + 1}`,
      style: 'width: 300px;',
    });
  }

  const resRow = {};

  answer.value?.result.map((row, i) => {
    resRow[i] = row;
  });

  rowsAnswer.value.push(resRow);
});
</script>

<style lang="scss" scoped></style>
