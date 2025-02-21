<template>
  <div class="q-pa-md flex column g-xl">
    <number-matrix
      :matrix="matrix"
      :columns-labels="labels"
      :matrix-title="matrixTitle"
    />
    <vectors-matrix
      v-model="matrixInputs"
      :is-done="isDone"
      @confirm="sendResult"
    />
    <vectors-matrix-answer
      v-if="isDone && correctAnswer && mark !== null"
      :matrix="correctAnswer"
      :max-mark="maxMark"
      :mark="mark"
    />
  </div>
</template>

<script lang="ts" setup>
import NumberMatrix from 'src/components/labs/NumberMatrix.vue';
import { reactive, ref, computed, watch } from 'vue';
import { getMatrixInputs } from 'src/lab-services/lab1a/helpers/get-matrix-input';
import VectorsMatrix from 'src/components/labs/VectorsMatrix.vue';
import { getMatrixCorrectPercent } from 'src/lab-services/custom/get-matrix-correct-percent';
import { solveMatrix } from 'src/lab-services/custom/solve-matrix';
import { getMarkForStep } from '../../../lab-services/custom/get-mark-for-matrix';
import { useLab1aStore } from 'src/stores/lab1a';
import VectorsMatrixAnswer from 'src/components/labs/VectorsMatrixAnswer.vue';
import { Lab1a } from 'src/models/lab/lab1a';

const props = defineProps<{
  matrix: number[][];
  maxMark: number;
  step: number;
  matrixTitle: string;
  labels: string[];
}>();

const labStore = useLab1aStore();

const matrixInputs = reactive<Lab1a.MatrixAnswer>(
  JSON.parse(
    localStorage.getItem(`user-answers-${props.step}`) ??
      JSON.stringify(getMatrixInputs(props.matrix))
  ) as Lab1a.MatrixAnswer
);

watch(matrixInputs, () => {
  localStorage.setItem(
    `user-answers-${props.step}`,
    JSON.stringify(matrixInputs)
  );
});

const isDone = computed(() =>
  info.value ? props.step < info.value?.step : false
);

const mark = ref<number | null>(
  JSON.parse(localStorage.getItem(`step-${props.step}-mark`) ?? 'null')
);

const info = computed(() => labStore.info);

const correctAnswer = computed<Lab1a.MatrixAnswer | null>(() =>
  isDone.value ? solveMatrix(props.matrix) : null
);

const sendResult = async () => {
  const correctAnswer = solveMatrix(props.matrix);
  const percent = getMatrixCorrectPercent(matrixInputs, correctAnswer);
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
