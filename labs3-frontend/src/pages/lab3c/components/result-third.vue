<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary">
      Введите оценку альтернативы a1 по критериям с помощью треугольных чисел
    </div>
    <q-table
      flat
      bordered
      title="Ответ"
      :rows="answers0[0]"
      :columns="columnsTemp0"
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
              v-if="col.name !== 'name' && currentStep === 0"
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
      v-if="currentStep === 0"
      label="Отправить"
      color="primary"
      @click="sendResult0"
    />
    <q-table
      v-if="currentStep !== 0"
      flat
      bordered
      title="Правильный ответ"
      :rows="answersToPrint0"
      :columns="columnsTemp0"
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
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div v-if="currentStep !== 0">
      Количество баллов за шаг: {{ finalAnswers0?.percentage }} из
      {{ finalAnswers0?.max_mark }}
    </div>
    <div class="text-h6 text-primary" v-if="currentStep !== 0">
      Введите нечеткую оценку альтернативы а1
    </div>
    <q-table
      v-if="currentStep !== 0"
      flat
      bordered
      title="Ответ"
      :rows="answers1"
      :columns="columnsTemp1"
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
              v-if="col.name !== 'name' && currentStep === 1"
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
      v-if="currentStep === 1"
      label="Отправить"
      color="primary"
      @click="sendResult1"
    />

    <q-table
      v-if="currentStep > 1"
      flat
      bordered
      title="Правильный ответ"
      :rows="answersToPrint1"
      :columns="columnsTemp1"
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
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div v-if="currentStep > 1">
      Количество баллов за шаг: {{ finalAnswers1?.percentage }} из
      {{ finalAnswers1?.max_mark }}
    </div>
    <div class="text-h6 text-primary" v-if="currentStep >= 2">
      Введите площадь фигуры под кривой треугольного числа
    </div>
    <q-input
      v-if="currentStep >= 2"
      v-model.number="answers2"
      @blur="answers2 = roundToNDigits(answers2)"
      style="max-width: 300px"
      :disable="currentStep !== 2"
      label="Площадь"
    />
    <q-btn
      v-if="currentStep === 2"
      label="Отправить"
      color="primary"
      @click="sendResult2"
    />
    <div v-if="currentStep > 2" class="text-h6 text-primary">
      Правильная площадь:
      <span :class="{ 'error-field': answersToPrint2?.flag === false }">{{
        answersToPrint2?.data
      }}</span>
    </div>
    <div v-if="currentStep > 2">
      Количество баллов за шаг: {{ finalAnswers2?.percentage }} из
      {{ finalAnswers2?.max_mark }}
    </div>
    <div class="text-h6 text-primary" v-if="currentStep >= 3">
      Определите коэффициенты уравнения прямой, на которой лежит большая из
      сторон треугольника
    </div>
    <div class="text-h6 text-primary flex column g-m" v-if="currentStep >= 3">
      <div>
        {{
          `a1: Y = ${answers3.k}x ${answers3.b > 0 ? '+' : ' '} ${answers3.b}`
        }}
      </div>
      <q-input
        v-model.number="answers3.k"
        @blur="answers3.k = roundToNDigits(answers3.k)"
        label="k"
        :disable="currentStep !== 3"
      />
      <q-input
        v-model.number="answers3.b"
        @blur="answers3.b = roundToNDigits(answers3.b)"
        label="b"
        :disable="currentStep !== 3"
      />
    </div>
    <q-btn
      v-if="currentStep === 3"
      label="Отправить"
      color="primary"
      @click="sendResult3"
    />
    <div v-if="currentStep > 3" class="text-h6 text-primary flex column g-m">
      <div>
        Правильный ответ: a1: Y =
        <span :class="{ 'error-field': answersToPrint3.k?.flag === false }">{{
          answersToPrint3.k.data
        }}</span
        >x {{ `${answersToPrint3.b.data > 0 ? '+' : ' '}` }}
        <span :class="{ 'error-field': answersToPrint3.b?.flag === false }">{{
          ` ${answersToPrint3.b.data}`
        }}</span>
      </div>
    </div>
    <div v-if="currentStep > 3">
      Количество баллов за шаг: {{ finalAnswers3?.percentage }} из
      {{ finalAnswers3?.max_mark }}
    </div>
    <div class="text-h6 text-primary" v-if="currentStep >= 4">
      Найдите координату по оси абсцисс, проходя через которую, прямая
      параллельная оси ординат делит треугольник, построенный по треугольному
      числу итоговой оценки альтернативы, пополам
    </div>
    <div v-if="currentStep >= 4" class="text-h6 text-primary flex column g-m">
      <div>
        {{ `a1: ${answers4.a1}x` }}<sup>2</sup>
        {{
          `${answers4.a2 > 0 ? '+' : '-'} ${Math.abs(answers4.a2)}x ${
            answers4.a3 > 0 ? '+' : '-'
          } ${Math.abs(answers4.a3)} = 0`
        }}
      </div>
      <q-input
        v-model.number="answers4.a1"
        @blur="answers4.a1 = roundToNDigits(answers4.a1)"
        label="a1"
        :disable="currentStep !== 4"
      />
      <q-input
        v-model.number="answers4.a2"
        @blur="answers4.a2 = roundToNDigits(answers4.a2)"
        label="a2"
        :disable="currentStep !== 4"
      />
      <q-input
        v-model.number="answers4.a3"
        @blur="answers4.a3 = roundToNDigits(answers4.a3)"
        label="a3"
        :disable="currentStep !== 4"
      />
    </div>
    <q-btn
      v-if="currentStep === 4"
      label="Отправить"
      color="primary"
      @click="sendResult4"
    />
    <div v-if="currentStep > 4" class="text-h6 text-primary flex column g-m">
      <div>
        Правильный ответ: a1:
        <span :class="{ 'error-field': answersToPrint4.a1?.flag === false }">{{
          answersToPrint4.a1.data
        }}</span
        >x<sup>2</sup> {{ `${answersToPrint4.a2.data > 0 ? '+' : '-'} `
        }}<span
          :class="{ 'error-field': answersToPrint4.a2?.flag === false }"
          >{{ Math.abs(answersToPrint4.a2.data) }}</span
        >x{{ `${answersToPrint4.a3.data > 0 ? '+' : '-'} `
        }}<span
          :class="{ 'error-field': answersToPrint4.a2?.flag === false }"
          >{{ Math.abs(answersToPrint4.a3.data) }}</span
        >
        = 0
      </div>
    </div>
    <div v-if="currentStep > 4">
      Количество баллов за шаг: {{ finalAnswers4?.percentage }} из
      {{ finalAnswers4?.max_mark }}
    </div>
    <div class="text-h6 text-primary" v-if="totalSum.maxMark">
      Общее количество полученных баллов за шаги: {{ totalSum.percentege }} из
      {{ totalSum.maxMark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';
import { Ref, computed, onMounted, reactive, ref, watch } from 'vue';
import { useLab3cStore } from 'src/stores/lab3c';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLab3cStore();

const currentStep = ref(0);

const totalSum = computed(() => {
  let maxMark = 0;
  let percentege = 0;

  if (finalAnswers0.value?.max_mark) {
    maxMark = finalAnswers0.value?.max_mark + maxMark;
    percentege = finalAnswers0.value?.percentage + percentege;
  }
  if (finalAnswers1.value?.max_mark) {
    maxMark = finalAnswers1.value?.max_mark + maxMark;
    percentege = finalAnswers1.value?.percentage + percentege;
  }
  if (finalAnswers2.value?.max_mark) {
    maxMark = finalAnswers2.value?.max_mark + maxMark;
    percentege = finalAnswers2.value?.percentage + percentege;
  }
  if (finalAnswers3.value?.max_mark) {
    maxMark = finalAnswers3.value?.max_mark + maxMark;
    percentege = finalAnswers3.value?.percentage + percentege;
  }
  if (finalAnswers4.value?.max_mark) {
    maxMark = finalAnswers4.value?.max_mark + maxMark;
    percentege = finalAnswers4.value?.percentage + percentege;
  }

  localStorage.setItem(
    'answer-3',
    JSON.stringify({ max_mark: maxMark, percentage: percentege })
  );

  return { maxMark, percentege };
});

const columnsTemp0 = ref([
  {
    name: 'name',
    field: 'name',
    label: '',
  },
  {
    name: '0',
    field: '0',
    label: '0',
    style: 'width: 300px',
  },
  {
    name: '1',
    field: '1',
    label: '1',
    style: 'width: 300px',
  },
  {
    name: '2',
    field: '2',
    label: '0',
    style: 'width: 300px',
  },
]);

const rowTemp0 = ref({
  name: 'В',
  '0': 0,
  '1': 0,
  '2': 0,
});

const rowNames = ['С1', 'С2', 'С3', 'С4'];

const answers0 = ref([]);

const finalAnswers0 = ref([]);

const answersToPrint0 = ref([]);

const columnsTemp1 = ref([
  {
    name: '0',
    field: '0',
    label: '0',
    style: 'width: 300px',
  },
  {
    name: '1',
    field: '1',
    label: '1',
    style: 'width: 300px',
  },
  {
    name: '2',
    field: '2',
    label: '0',
    style: 'width: 300px',
  },
]);

const answers1 = ref([]);

const finalAnswers1 = ref([]);

const answersToPrint1 = ref([]);

const answers2 = ref(0);

const finalAnswers2 = ref(0);

const answersToPrint2 = ref(0);

const answers3 = ref({
  k: 1,
  b: 1,
});

const finalAnswers3 = ref(0);

const answersToPrint3 = ref(0);

const answers4 = ref({
  a1: 1,
  a2: 1,
  a3: 1,
});

const finalAnswers4 = ref(0);

const answersToPrint4 = ref(0);

onMounted(() => {
  const localCurrentStep = localStorage.getItem('current-step-3');

  if (localCurrentStep) {
    currentStep.value = parseInt(localCurrentStep);
  }

  const localAnswers0 = localStorage.getItem('user-answers-3-0');
  const localFinalAnswers0 = localStorage.getItem('answer-3-0');

  if (localAnswers0) {
    answers0.value = JSON.parse(localAnswers0);
  } else {
    const answer0 = [];
    props.variant.criteria.forEach((_, index) => {
      answer0.push({
        ...JSON.parse(JSON.stringify(rowTemp0.value)),
        name: rowNames[index],
      });
    });
    answers0.value.push(answer0);
  }

  if (localFinalAnswers0) {
    finalAnswers0.value = JSON.parse(localFinalAnswers0);
  }

  const localAnswers1 = localStorage.getItem('user-answers-3-1');
  const localFinalAnswers1 = localStorage.getItem('answer-3-1');

  if (localAnswers1) {
    answers1.value = JSON.parse(localAnswers1);
  } else {
    answers1.value.push({
      ...JSON.parse(JSON.stringify(rowTemp0.value)),
    });
  }

  if (localFinalAnswers1) {
    finalAnswers1.value = JSON.parse(localFinalAnswers1);
  }

  const localAnswers2 = localStorage.getItem('user-answers-3-2');
  const localFinalAnswers2 = localStorage.getItem('answer-3-2');

  if (localAnswers2) {
    answers2.value = JSON.parse(localAnswers2);
  } else {
    answers2.value = 0;
  }

  if (localFinalAnswers2) {
    finalAnswers2.value = JSON.parse(localFinalAnswers2);
  }

  const localAnswers3 = localStorage.getItem('user-answers-3-3');
  const localFinalAnswers3 = localStorage.getItem('answer-3-3');

  if (localAnswers3) {
    answers3.value = JSON.parse(localAnswers3);
  } else {
    answers3.value = {
      k: 1,
      b: 1,
    };
  }

  if (localFinalAnswers3) {
    finalAnswers3.value = JSON.parse(localFinalAnswers3);
  }

  const localAnswers4 = localStorage.getItem('user-answers-3-4');
  const localFinalAnswers4 = localStorage.getItem('answer-3-4');

  if (localAnswers4) {
    answers4.value = JSON.parse(localAnswers4);
  } else {
    answers4.value = {
      a1: 1,
      a2: 1,
      a3: 1,
    };
  }

  if (localFinalAnswers4) {
    finalAnswers4.value = JSON.parse(localFinalAnswers4);
  }
});

watch(
  answers0,
  () => {
    localStorage.setItem('user-answers-3-0', JSON.stringify(answers0.value));
  },
  { deep: true }
);

watch(
  currentStep,
  () => {
    localStorage.setItem('current-step-3', JSON.stringify(currentStep.value));
  },
  { deep: true }
);

watch(
  finalAnswers0.value,
  () => {
    localStorage.setItem('answer-3-0', JSON.stringify(finalAnswers0.value));
  },
  { deep: true }
);

const sendResult0 = async () => {
  const res = answers0.value[0].map((a) => {
    const resRow = [];
    Object.keys(a).forEach((key) => {
      if (key !== 'name') {
        const k = key == '2' ? '0' : key;
        const o = {
          X: roundToNDigits(a[key]),
          Y: roundToNDigits(k),
        };
        resRow.push(o);
      }
    });
    return resRow;
  });

  const resAnswer = await labStore.sendAlternativeMatrices({
    matrix: res,
    step: 1,
  });

  if (resAnswer.data) {
    finalAnswers0.value = resAnswer.data;
    currentStep.value = 1;
    localStorage.setItem('answer-3-0', JSON.stringify(finalAnswers0.value));
    labStore.getInfo();
  }
};

watch(
  finalAnswers0,
  () => {
    if (finalAnswers0.value?.result?.length) {
      answersToPrint0.value = finalAnswers0.value.result.map((ans, index) => {
        const resO = { name: { data: `C${index + 1}` } };
        ans.forEach((o, i) => {
          resO[i] = { data: o.X, flag: o.Flag };
        });
        return resO;
      });
    }
  },
  { deep: true }
);

onMounted(() => {
  if (!finalAnswers0.value?.result?.length) return;
  answersToPrint0.value = finalAnswers0.value.result.map((ans, index) => {
    const resO = { name: { data: `C${index + 1}` } };
    ans.forEach((o, i) => {
      resO[i] = { data: o.X, flag: o.Flag };
    });
    return resO;
  });
});

watch(
  answers1,
  () => {
    localStorage.setItem('user-answers-3-1', JSON.stringify(answers1.value));
  },
  { deep: true }
);

watch(
  finalAnswers1,
  () => {
    localStorage.setItem('answer-3-1', JSON.stringify(finalAnswers1.value));
  },
  { deep: true }
);

const sendResult1 = async () => {
  const res = answers1.value.map((a) => {
    const resRow = [];
    Object.keys(a).forEach((key) => {
      if (key !== 'name') {
        const k = key == '2' ? '0' : key;
        const o = {
          X: roundToNDigits(a[key]),
          Y: roundToNDigits(k),
        };
        resRow.push(o);
      }
    });
    return resRow;
  });

  const resAnswer = await labStore.sendCriteriaEstimation({
    matrix: res[0],
    step: 1,
  });

  if (resAnswer.data) {
    finalAnswers1.value = resAnswer.data;
    currentStep.value = 2;
    localStorage.setItem('answer-3-1', JSON.stringify(finalAnswers1.value));
    labStore.getInfo();
  }
};

watch(
  finalAnswers1,
  () => {
    if (finalAnswers1.value?.result?.length) {
      const resO = {};

      finalAnswers1.value.result.forEach((o, i) => {
        resO[i] = { data: o.X, flag: o.Flag };
      });

      answersToPrint1.value = [resO];
    }
  },
  {
    deep: true,
  }
);

onMounted(() => {
  if (!finalAnswers1.value?.result?.length) return;
  const resO = {};

  finalAnswers1.value.result.forEach((o, i) => {
    resO[i] = { data: o.X, flag: o.Flag };
  });

  answersToPrint1.value = [resO];
});

watch(
  answers2,
  () => {
    localStorage.setItem('user-answers-3-2', JSON.stringify(answers2.value));
  },
  { deep: true }
);

watch(
  finalAnswers2,
  () => {
    localStorage.setItem('answer-3-2', JSON.stringify(finalAnswers2.value));
  },
  { deep: true }
);

const sendResult2 = async () => {
  const resAnswer = await labStore.sendArea({
    set: roundToNDigits(answers2.value),
    step: 1,
  });

  if (resAnswer.data) {
    finalAnswers2.value = resAnswer.data;
    currentStep.value = 3;
    localStorage.setItem('answer-3-2', JSON.stringify(finalAnswers2.value));
    labStore.getInfo();
  }
};

watch(
  finalAnswers2,
  () => {
    if (finalAnswers2.value) {
      answersToPrint2.value = finalAnswers2.value.result;
    }
  },
  {
    deep: true,
  }
);

onMounted(() => {
  if (!finalAnswers2.value) return;
  answersToPrint2.value = finalAnswers2.value.result;
});

watch(
  answers3,
  () => {
    localStorage.setItem('user-answers-3-3', JSON.stringify(answers3.value));
  },
  { deep: true }
);

watch(
  finalAnswers3,
  () => {
    localStorage.setItem('answer-3-3', JSON.stringify(finalAnswers3.value));
  },
  { deep: true }
);

const sendResult3 = async () => {
  const resAnswer = await labStore.sendLine({
    parameters: {
      k: roundToNDigits(answers3.value.k),
      b: roundToNDigits(answers3.value.b),
    },
    step: 1,
  });

  if (resAnswer.data) {
    finalAnswers3.value = resAnswer.data;
    currentStep.value = 4;
    localStorage.setItem('answer-3-3', JSON.stringify(finalAnswers3.value));
    labStore.getInfo();
  }
};

watch(finalAnswers3, () => {
  if (finalAnswers3.value) {
    answersToPrint3.value = finalAnswers3.value.result.parameters;
  }
});

onMounted(() => {
  if (!finalAnswers3.value) return;
  answersToPrint3.value = finalAnswers3.value.result.parameters;
});

watch(
  answers4,
  () => {
    localStorage.setItem('user-answers-3-4', JSON.stringify(answers4.value));
  },
  { deep: true }
);

watch(
  finalAnswers4,
  () => {
    localStorage.setItem('answer-3-4', JSON.stringify(finalAnswers4.value));
  },
  { deep: true }
);

const sendResult4 = async () => {
  const resAnswer = await labStore.sendQuadratic({
    parameters: {
      a1: roundToNDigits(answers4.value.a1),
      a2: roundToNDigits(answers4.value.a2),
      a3: roundToNDigits(answers4.value.a3),
    },
    step: 1,
  });

  if (resAnswer.data) {
    finalAnswers4.value = resAnswer.data;
    currentStep.value = 5;
    localStorage.setItem('answer-3-4', JSON.stringify(finalAnswers4.value));
    const step4 = localStorage.getItem('current-step-4');
    const step5 = localStorage.getItem('current-step-5');
    if (step4 === '5' && step5 === '5') {
      await labStore.increment2Step();
    }
    labStore.normalStep = 4;
    labStore.getInfo();
  }
};

watch(finalAnswers4, () => {
  if (finalAnswers4.value) {
    answersToPrint4.value = finalAnswers4.value.result.parameters;
  }
});

onMounted(() => {
  if (!finalAnswers4.value) return;
  answersToPrint4.value = finalAnswers4.value.result.parameters;
});
</script>

<style lang="scss" scoped></style>
