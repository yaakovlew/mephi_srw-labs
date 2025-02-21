<template>
  <div class="q-pa-md flex column g-m">
    <div class="text-h6 text-primary text-center text-bold">
      2 - Лабораторная работа "Групповая экспертиза"
    </div>
    <div
      class="page-title text-primary flex items-center"
      style="position: absolute"
    >
      <countdown ref="countdownRef" :time="900000 * 4" />
    </div>
    <div
      class="page-title text-primary flex items-center g-xl full-width justify-center text-center"
    >
      Получено баллов: {{ info?.percentage }} / 100
      <br />
      Текущий этап: {{ (info?.step ?? 1) - 1 }} / 6
    </div>
    <q-tabs v-model="tab" inline-label shrink stretch class="text-primary">
      <q-tab v-for="tab in tabs" :key="tab.name" v-bind="tab" />
    </q-tabs>
    <q-tab-panels
      v-if="variant"
      v-model="tab"
      animated
      class="shadow-2 rounded-borders"
    >
      <q-tab-panel name="task">
        <task-description :variant="variant" :variant-number="24" />
      </q-tab-panel>
      <q-tab-panel name="result-1">
        <mean-values
          :table="variant.tables[0]"
          :matrix-title="variant.description[0]"
          :max-mark="10"
          :step="1"
        />
      </q-tab-panel>
      <q-tab-panel name="result-2">
        <weights-values
          :table="variant.tables[0]"
          :matrix-title="variant.description[0]"
          :max-mark="10"
          :step="2"
        />
      </q-tab-panel>
      <q-tab-panel name="result-3">
        <bayes-values
          :table="variant.tables[0]"
          :matrix-title="variant.description[0]"
          :max-mark="10"
          :step="3"
        />
      </q-tab-panel>
      <q-tab-panel name="result-4">
        <mean-values-short
          :table="variant.tables[1]"
          :matrix-title="variant.description[1]"
          :max-mark="10"
          :step="4"
        />
      </q-tab-panel>
      <q-tab-panel name="result-5">
        <weights-values
          :table="variant.tables[1]"
          :matrix-title="variant.description[1]"
          :max-mark="10"
          :step="5"
        />
      </q-tab-panel>
      <q-tab-panel name="result-6">
        <bayes-values-short
          :table="variant.tables[1]"
          :matrix-title="variant.description[1]"
          :max-mark="10"
          :step="6"
        />
      </q-tab-panel>
      <q-tab-panel name="result-7">
        <weights-difference
          :table-first="variant.tables[1]"
          :table-second="variant.tables[0]"
          :variant="variant"
          matrix-title="Веса"
          :max-mark="10"
          :step="7"
        />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
// import TaskDescription from './components/task-description.vue';
import Countdown from 'vue3-countdown';
import { useLab2Store } from 'src/stores/lab2';
import { getVariant } from 'src/lab-services/lab2/helpers/get-variant';
import { getMaxMarksForStep } from 'src/lab-services/lab1a/lab1a';
import bayesValues from './components/bayes-values.vue';
import { Notify } from 'quasar';
import meanValues from './components/mean-values.vue';
import weightsValues from './components/weights-values.vue';
import meanValuesShort from './components/mean-values-short.vue';
import bayesValuesShort from './components/bayes-values-short.vue';
import weightsDifference from './components/weights-difference.vue';

const labStore = useLab2Store();

localStorage.setItem(
  'lab-token',
  'sf4lmlskmkdflkmsdlkfmskldmflsdkmfskldfmksdlkfm'
);

const maxMarks = getMaxMarksForStep();

const variant = computed(() => labStore.variant);

const currentStep = computed(() => labStore.info?.step);

const info = computed(() => labStore.info);

const countdownRef = ref();

const pause = () => countdownRef.value.stop();

watch(currentStep, () => {
  tabs.value = fullSteps.filter(
    (tab, i) =>
      (currentStep.value !== undefined && i <= currentStep.value) ||
      tab.name === 'result'
  );
  if (currentStep.value === 8) {
    pause();
    Notify.create({
      type: 'positive',
      message: `Вы сделали лабораторную работу и получили ${info.value?.percentage} баллов из 100`,
    });
    labStore.sendResult();
  }
});

const fullSteps = [
  {
    name: 'task',
    icon: 'info',
    label: 'Задание',
  },
  {
    name: 'result-1',
    icon: 'done',
    label: 'Этап 1',
  },
  {
    name: 'result-2',
    icon: 'done',
    label: 'Этап 2',
  },
  {
    name: 'result-3',
    icon: 'done',
    label: 'Этап 3',
  },
  {
    name: 'result-4',
    icon: 'done',
    label: 'Этап 4',
  },
  {
    name: 'result-5',
    icon: 'done',
    label: 'Этап 5',
  },
  {
    name: 'result-6',
    icon: 'done',
    label: 'Этап 6',
  },
  {
    name: 'result-7',
    icon: 'done',
    label: 'Этап 7',
  },
];

onMounted(async () => {
  await labStore.getInfo();
  const v = getVariant(7);
  await labStore.saveVariant({ number: v.variantNumber, data: v.variant });
});

const tabs = ref([
  {
    name: 'task',
    icon: 'info',
    label: 'Задание',
  },
  {
    name: 'result-1',
    icon: 'done',
    label: 'Этап 1',
  },
]);

const tab = ref('task');
</script>

<style lang="scss" scoped></style>
