<template>
  <div class="q-pa-md flex column g-m">
    <div class="text-h6 text-primary text-center text-bold">
      3А - Лабораторная работа "Метод отношений предпочтений"
    </div>
    <div
      class="page-title text-primary flex items-center"
      style="position: absolute"
    >
		<countdown ref="countdownRef" :time="24 * 900000 * 4" @finish="onTimerFinish" />
    </div>
    <q-btn
      style="position: absolute; right: 5px"
      label="Об округлении"
      icon="info"
      color="primary"
      flat
      @click="alert = true"
    />
    <q-dialog v-model="alert">
      <q-card>
        <q-card-section class="text-h6 text-primary">
          Числа округляются до двух знаков после запятой
        </q-card-section>
      </q-card>
    </q-dialog>
    <div
      class="page-title text-primary flex items-center g-xl full-width justify-center text-center"
    >
      Получено баллов: {{ fullInfo?.percentage }} / 100
      <br />
      Текущий этап: {{ fullInfo?.step }} / 7
    </div>
    <q-tabs v-model="tab" inline-label shrink stretch class="text-primary">
      <q-tab v-for="tab in tabs" :key="tab.name" v-bind="tab" />
    </q-tabs>
    <q-tab-panels v-model="tab" animated class="shadow-2 rounded-borders">
      <q-tab-panel name="task">
        <task-description :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-1">
        <result-first :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-2">
        <result-second :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-3">
        <result-third :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-4">
        <result-fourth :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-5">
        <result-fifth :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-6">
        <result-sixth :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result-7">
        <result-seventh :variant="variant" />
      </q-tab-panel>
      <q-tab-panel name="result">
        <lab-result />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
import TaskDescription from './components/task-description.vue';
import ResultFirst from './components/result-first.vue';
import ResultSecond from './components/result-second.vue';
import ResultThird from './components/result-third.vue';
import ResultFourth from './components/result-fourth.vue';
import ResultFifth from './components/result-fifth.vue';
import ResultSixth from './components/result-sixth.vue';
import ResultSeventh from './components/result-seventh.vue';
import labResult from './components/lab-result.vue';
import { useLabStore } from 'src/stores/lab';
import Countdown from 'vue3-countdown';
import { Notify } from 'quasar';

const countdownRef = ref();

localStorage.setItem('lab-token', '38419u38uroefeniuhf9835f34ub20yb5b2y522ny5');

const pause = () => countdownRef.value.stop();

const labStore = useLabStore();

const userId = computed(() => labStore.userId);
const variant = computed(() => labStore.variant);

const currentStep = computed(() => labStore.info?.step);

const fullInfo = computed(() => labStore.info);

const alert = ref(false);

watch(currentStep, () => {
  tabs.value = fullSteps.filter(
    (tab, i) =>
      (currentStep.value !== undefined && i <= currentStep.value + 1) ||
      tab.name === 'result'
  );
  if (currentStep.value === 7) {
    pause();
    Notify.create({
      type: 'positive',
      message: `Вы сделали лабораторную работу и получили ${fullInfo.value?.percentage} баллов из 100`,
    });
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
  {
    name: 'result',
    icon: 'done',
    label: 'Результаты',
  },
];

onMounted(async () => {
  await labStore.getLabVariant();
  await labStore.getInfo();
});

const tabs = ref([
  {
    name: 'task',
    icon: 'info',
    label: 'Задание',
  },
]);

const tab = ref('task');
const panel = ref('task');

const onTimerFinish = () => {
  localStorage.setItem('is_done', '1');
  Notify.create({
    type: 'warning',
    message: 'Время выполнения лабораторной работы истекло!',
  });
};

window.onbeforeunload = closingCode;

function closingCode() {
  if (currentStep.value === 7) {
    localStorage.removeItem('user-answers-1');
    localStorage.removeItem('user-answers-2');
    localStorage.removeItem('user-answers-3');
    localStorage.removeItem('user-answers-4');
    localStorage.removeItem('user-answers-5');
    localStorage.removeItem('user-answers-6');
    localStorage.removeItem('user-answers-7');
    localStorage.removeItem('user-answers-1-1');
    localStorage.removeItem('user-answers-1-2');
    localStorage.removeItem('user-answers-1-3');
    localStorage.removeItem('user-answers-1-4');
    localStorage.removeItem('chosen-index-7');
    localStorage.removeItem('answer-1');
    localStorage.removeItem('answer-2');
    localStorage.removeItem('answer-3');
    localStorage.removeItem('answer-4');
    localStorage.removeItem('answer-5');
    localStorage.removeItem('answer-6');
    localStorage.removeItem('answer-7');
    localStorage.removeItem('answer-8');
    localStorage.removeItem('answer-9');
    localStorage.removeItem('answer-10');
    localStorage.removeItem('answer-2-1');
    localStorage.removeItem('answer-2-2');
    localStorage.removeItem('answer-2-3');
    localStorage.removeItem('answer-2-4');
		localStorage.removeItem('answer-2-5');
  }
  // do something...
  return null;
}
</script>

<style lang="scss" scoped></style>
