<template>
  <div class="q-pa-md flex column g-xl">
    <number-matrix
      :matrix="weights"
      :columns-labels="columnLabels"
      :rows-labels="rowLabels"
      :matrix-title="matrixTitle"
    />
    <div class="text-h5 text-primary">Ответ</div>
    <six-inputes v-model="answers.diff" :is-done="isDone" title="Разница" />
    <q-btn
      v-if="!isDone"
      flat
      color="primary"
      label="Отправить"
      @click="sendResult"
    />
    <div class="text-h5 text-primary" v-if="isDone">Правильный ответ</div>
    <six-inputes-answer
      v-model="correctAnswers.diff"
      :is-done="isDone"
      title="Разница"
    />
    <div v-if="isDone" class="flex column g-m text-h6 text-primary">
      Баллов за этап {{ mark }} из {{ maxMark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab2 } from 'src/models/lab/lab2';
import NumberMatrix from 'src/components/labs/NumberMatrix.vue';
import sixInputes from '../base/six-inputes.vue';
import sixInputesAnswer from '../base/six-inputes-answer.vue';
import { computed, watch, onMounted } from 'vue';
import { ref } from 'vue';
import { useLab2Store } from '../../../stores/lab2';
import { reactive } from 'vue';
import { getMarkForStep } from 'src/lab-services/custom/get-mark-for-matrix';
import { getMatrixCorrectPercent } from 'src/lab-services/custom/get-matrix-correct-percent';
import { getWeightsDifference } from 'src/lab-services/lab2/lab2';
import { getMatrixWeightVector } from 'src/lab-services/lab2/calculators';

const props = defineProps<{
  tableFirst: Lab2.Table;
  tableSecond: Lab2.Table;
  matrixTitle: string;
  step: number;
  maxMark: number;
  variant: Lab2.Variant;
}>();

const firstWeights = getMatrixWeightVector(props.tableFirst);
const correctWeights = getMatrixWeightVector(props.tableSecond);

const weights = [firstWeights, correctWeights];

const labStore = useLab2Store();

const columnLabels = [
  'Критерий 1',
  'Критерий 2',
  'Критерий 3',
  'Критерий 4',
  'Критерий 5',
  'Критерий 6',
];

const rowLabels = ['Выборка 1', 'Выборка 2'];

const info = computed(() => labStore.info);

const isDone = computed(() =>
  info.value ? props.step < info.value?.step : false
);

const answers = reactive(
  JSON.parse(
    localStorage.getItem(`user-answers-${props.step}`) ??
      '{"diff":[0,0,0,0,0,0]}'
  )
);

const correctAnswers = reactive({
  diff: [0, 0, 0, 0, 0, 0],
});

const mark = ref<number | null>(
  JSON.parse(localStorage.getItem(`step-${props.step}-mark`) ?? 'null')
);

watch(answers, () => {
  localStorage.setItem(`user-answers-${props.step}`, JSON.stringify(answers));
});

onMounted(() => {
  if (isDone.value) {
    correctAnswers.diff = getWeightsDifference(props.variant);
  }
});

const sendResult = async () => {
  correctAnswers.diff = getWeightsDifference(props.variant);

  const percent = getMatrixCorrectPercent(answers, correctAnswers);
  mark.value = getMarkForStep(percent, props.maxMark);

  localStorage.setItem(`step-${props.step}-mark`, JSON.stringify(mark.value));
  if (mark.value !== null && info.value) {
    await labStore.updateInfo({
      step: props.step + 1,
      percentage: mark.value,
    });
    await labStore.getInfo();
  }
};
</script>

<style lang="scss" scoped></style>
