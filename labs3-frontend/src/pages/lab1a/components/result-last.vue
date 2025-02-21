<template>
  <div class="q-pa-md flex column g-xl">
    <number-matrix
      :matrix="matrix"
      :columns-labels="labels"
      :matrix-title="matrixTitle"
      :rows-labels="rowLabels"
    />
    <priorities-alternatives
      v-model="prioritiesAlternatives"
      :alternatives="alternatives"
      :is-done="isDone"
      :criterias="criterias"
      :weights="weights"
      @confirm="sendResult"
    />
    <priorities-alternatives-answer
      v-if="correctAnswer && mark !== null"
      :correct-answer="correctAnswer"
      :alternatives="alternatives"
      :max-mark="maxMark"
      :mark="mark"
    />
  </div>
</template>

<script lang="ts" setup>
import NumberMatrix from 'src/components/labs/NumberMatrix.vue';
import { Lab1a } from 'src/models/lab/lab1a';
import { getFinalMatrix, solveVariant } from 'src/lab-services/lab1a/lab1a';
import PrioritiesAlternatives from 'src/components/labs/PrioritiesAlternatives.vue';
import PrioritiesAlternativesAnswer from 'src/components/labs/PrioritiesAlternativesAnswer.vue';
import { useLab1aStore } from 'src/stores/lab1a';
import { computed, reactive, ref, watch } from 'vue';
import { getAlternativesInputs } from 'src/lab-services/lab1a/helpers/get-alternatives-inputs';
import { getResultCorrectPercent } from 'src/lab-services/custom/get-matrix-correct-percent';
import { getMarkForStep } from 'src/lab-services/custom/get-mark-for-matrix';

const props = defineProps<{
  variant: Lab1a.Variant;
  maxMark: number;
  step: number;
  matrixTitle: string;
  labels: string[];
  rowLabels: string[];
  alternatives: string[];
  criterias: string[];
}>();

const finalMatrix = getFinalMatrix(props.variant);

const matrix = finalMatrix.finalMatrix;

const weights = finalMatrix.finalWeights;

const labStore = useLab1aStore();

const info = computed(() => labStore.info);

const prioritiesAlternatives = reactive<Lab1a.PrioritiesAlternatives>(
  JSON.parse(
    localStorage.getItem(`user-answers-${props.step}`) ??
      JSON.stringify(getAlternativesInputs(props.alternatives))
  ) as Lab1a.PrioritiesAlternatives
);

watch(prioritiesAlternatives, () => {
  localStorage.setItem(
    `user-answers-${props.step}`,
    JSON.stringify(prioritiesAlternatives)
  );
});

const isDone = computed(() =>
  info.value ? props.step < info.value?.step : false
);

const mark = ref<number | null>(
  JSON.parse(localStorage.getItem(`step-${props.step}-mark`) ?? 'null')
);

const correctAnswer = computed<Lab1a.VariantAnswer | null>(() =>
  isDone.value ? solveVariant(props.variant) : null
);

const sendResult = async () => {
  const correctAnswer = solveVariant(props.variant);
  if (!correctAnswer) return;
  const percent = getResultCorrectPercent(
    prioritiesAlternatives,
    correctAnswer
  );
  mark.value = getMarkForStep(percent, props.maxMark);
  localStorage.setItem(`step-${props.step}-mark`, JSON.stringify(mark.value));
  if (mark.value !== null && info.value) {
    await labStore.updateInfo({
      step: props.step + 1,
      percentage: info.value?.percentage + mark.value,
    });
    await labStore.getInfo();
  }
};
</script>

<style lang="scss" scoped></style>
