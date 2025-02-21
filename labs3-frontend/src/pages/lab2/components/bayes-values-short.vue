<template>
  <div class="q-pa-md flex column g-xl">
    <number-matrix
      :matrix="table"
      :columns-labels="labels"
      :rows-labels="rowLabels"
      :matrix-title="matrixTitle"
    />
    <six-inputes-answer
      v-model="infoForBias.weights"
      :is-done="true"
      title="Веса"
    />
    <div class="text-h5 text-primary">Ответ</div>
    <six-inputes
      v-model="answers.bayes"
      :is-done="isDone"
      title="Среднее по Байесу"
    />
    <q-btn
      v-if="!isDone"
      flat
      color="primary"
      label="Отправить"
      @click="sendResult"
    />
    <div class="text-h5 text-primary" v-if="isDone">Правильный ответ</div>
    <six-inputes-answer
      v-model="correctAnswers.bayes"
      :is-done="isDone"
      title="Среднее по Байесу"
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
import { getMatrixWeightVector } from '../../../lab-services/custom/solve-matrix';
import { getMatrixBayesMean } from '../../../lab-services/lab2/calculators';

const props = defineProps<{
  table: Lab2.Table;
  matrixTitle: string;
  step: number;
  maxMark: number;
}>();

const labStore = useLab2Store();

const labels = [
  'Критерий 1',
  'Критерий 2',
  'Критерий 3',
  'Критерий 4',
  'Критерий 5',
  'Критерий 6',
];

const info = computed(() => labStore.info);

const isDone = computed(() =>
  info.value ? props.step < info.value?.step : false
);

const rowLabels = computed(() => props.table.map((_, i) => `Экспрет ${i + 1}`));

const answers = reactive(
  JSON.parse(
    localStorage.getItem(`user-answers-${props.step}`) ??
      '{"bayes":[0,0,0,0,0,0]}'
  )
);

const correctAnswers = reactive({
  bayes: [0, 0, 0, 0, 0, 0],
});

const infoForBias = reactive({
  weights: [0, 0, 0, 0, 0, 0],
});

const mark = ref<number | null>(
  JSON.parse(localStorage.getItem(`step-${props.step}-mark`) ?? 'null')
);

watch(answers, () => {
  localStorage.setItem(`user-answers-${props.step}`, JSON.stringify(answers));
});

onMounted(() => {
  if (isDone.value) {
    correctAnswers.bayes = getMatrixBayesMean(props.table);
  }
  infoForBias.weights = getMatrixWeightVector(props.table);
});

const sendResult = async () => {
  correctAnswers.bayes = getMatrixBayesMean(props.table);

  const percent = getMatrixCorrectPercent(answers, correctAnswers);
  mark.value = getMarkForStep(percent, props.maxMark);

  console.log(percent, mark.value, answers, correctAnswers);
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
