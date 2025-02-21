<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary">
      Постройте соответствующие отношения предпочтений по каждому из критериев
    </div>
    <span class="text-h6 text-primary">C1</span>
    <q-table
      flat
      bordered
      title="Ответ"
      :rows="rows1"
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
              v-if="col.name !== 'name' && stepStep === 1 && !isDone"
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
      v-if="stepStep === 1 && !isDone"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="stepStep !== 1"
      flat
      bordered
      title="Правильный ответ"
      :rows="rows1Answer"
      :columns="columnsAnswer"
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
    <div v-if="answer1">
      Количество полученных баллов за шаг: {{ answer1.percentage }} из
      {{ answer1.max_mark }}
    </div>
    <span v-if="stepStep !== 1" class="text-h6 text-primary">C2</span>
    <q-table
      v-if="stepStep !== 1"
      flat
      bordered
      title="Ответ"
      :rows="rows2"
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
              v-if="col.name !== 'name' && stepStep === 2 && !isDone"
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
      v-if="stepStep === 2 && !isDone"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="stepStep !== 2 && stepStep !== 1"
      flat
      bordered
      title="Правильный ответ"
      :rows="rows2Answer"
      :columns="columnsAnswer"
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
    <div v-if="answer2">
      Количество полученных баллов за шаг: {{ answer2.percentage }} из
      {{ answer2.max_mark }}
    </div>
    <span class="text-h6 text-primary" v-if="stepStep !== 2 && stepStep !== 1"
      >C3</span
    >
    <q-table
      v-if="stepStep !== 2 && stepStep !== 1"
      flat
      bordered
      title="Ответ"
      :rows="rows3"
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
              v-if="col.name !== 'name' && stepStep === 3 && !isDone"
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
      v-if="stepStep === 3 && !isDone"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="stepStep === 4 || stepStep === 5"
      flat
      bordered
      title="Правильный ответ"
      :rows="rows3Answer"
      :columns="columnsAnswer"
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
    <div v-if="answer3">
      Количество полученных баллов за шаг: {{ answer3.percentage }} из
      {{ answer3.max_mark }}
    </div>
    <span class="text-h6 text-primary" v-if="stepStep === 4 || stepStep === 5"
      >C4</span
    >
    <q-table
      v-if="stepStep === 4 || stepStep === 5"
      flat
      bordered
      title="Ответ"
      :rows="rows4"
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
              v-if="col.name !== 'name' && stepStep === 4 && !isDone"
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
      v-if="stepStep === 4 && !isDone"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <q-table
      v-if="stepStep === 5"
      flat
      bordered
      title="Правильный ответ"
      :rows="rows4Answer"
      :columns="columnsAnswer"
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
    <div v-if="answer4">
      Количество полученных баллов за шаг: {{ answer4.percentage }} из
      {{ answer4.max_mark }}
    </div>

    <div v-if="answer" style="font-size: 18px">
      Общее количество полученных баллов за шаги: {{ totalSum.sum }} из
      {{ totalSum.maxCount }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useLabStore } from '../../../stores/lab';
import { data } from 'autoprefixer';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLabStore();

const columns = ref([]);
const rows1 = ref([]);

const rows2 = ref([]);

const rows3 = ref([]);

const rows4 = ref([]);

const isDone = computed(() => labStore.info?.step !== 1);

const stepStep = ref(1);

onMounted(() => {
  if (!props.variant) return;

  const userAnswers1 = localStorage.getItem('user-answers-1-1');
  if (userAnswers1) {
    rows1.value = JSON.parse(userAnswers1);
  }
  const userAnswers2 = localStorage.getItem('user-answers-1-2');
  if (userAnswers2) {
    rows2.value = JSON.parse(userAnswers2);
  }
  const userAnswers3 = localStorage.getItem('user-answers-1-3');
  if (userAnswers3) {
    rows3.value = JSON.parse(userAnswers3);
  }
  const userAnswers4 = localStorage.getItem('user-answers-1-4');
  if (userAnswers4) {
    rows4.value = JSON.parse(userAnswers4);
  }
  columns.value.push({
    name: 'name',
    required: true,
    label: `Название`,
  });

  for (let i = 0; i < 3; i++) {
    columns.value.push({
      name: i,
      required: true,
      label: `A${i + 1}`,
      style: 'width: 300px;',
    });
    if (!userAnswers1) {
      rows1.value.push({
        name: `A${i + 1}`,
        0: 0,
        1: 0,
        2: 0,
      });
    }
    if (!userAnswers2) {
      rows2.value.push({
        name: `A${i + 1}`,
        0: 0,
        1: 0,
        2: 0,
      });
    }

    if (!userAnswers3) {
      rows3.value.push({
        name: `A${i + 1}`,
        0: 0,
        1: 0,
        2: 0,
      });
    }
    if (!userAnswers4) {
      rows4.value.push({
        name: `A${i + 1}`,
        0: 0,
        1: 0,
        2: 0,
      });
    }
  }
});

const totalSum = computed(() => {
  let sum = 0;
  let maxCount = 0;
  answer.value.forEach((item) => {
    if (item) {
      sum += item.percentage;
      maxCount += item.max_mark;
    }
  });

  localStorage.setItem(
    'answer-2',
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

watch(
  rows1,
  () => {
    localStorage.setItem('user-answers-1-1', JSON.stringify(rows1.value));
  },
  {
    deep: true,
  }
);

watch(
  rows2,
  () => {
    localStorage.setItem('user-answers-1-2', JSON.stringify(rows2.value));
  },
  {
    deep: true,
  }
);

watch(
  rows3,
  () => {
    localStorage.setItem('user-answers-1-3', JSON.stringify(rows3.value));
  },
  {
    deep: true,
  }
);

watch(
  rows4,
  () => {
    localStorage.setItem('user-answers-1-4', JSON.stringify(rows4.value));
  },
  {
    deep: true,
  }
);

const isSend = ref(false);
const answer1: Ref<Lab.AlternativeDiffMatricesAnswer | null> = ref(
  localStorage.getItem('answer-2-1')
    ? JSON.parse(localStorage.getItem('answer-2-1')!)
    : null
);
const answer2: Ref<Lab.AlternativeDiffMatricesAnswer | null> = ref(
  localStorage.getItem('answer-2-2')
    ? JSON.parse(localStorage.getItem('answer-2-2')!)
    : null
);
const answer3: Ref<Lab.AlternativeDiffMatricesAnswer | null> = ref(
  localStorage.getItem('answer-2-3')
    ? JSON.parse(localStorage.getItem('answer-2-3')!)
    : null
);
const answer4: Ref<Lab.AlternativeDiffMatricesAnswer | null> = ref(
  localStorage.getItem('answer-2-4')
    ? JSON.parse(localStorage.getItem('answer-2-4')!)
    : null
);

const sendResult = async () => {
  const res1 = rows1.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  const res2 = rows2.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  const res3 = rows3.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  const res4 = rows4.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a] === '' ? -1 : roundToNDigits(row[a]));
      }
    }
    return resRow;
  });

  let res = [];

  // const res = [res1, res2, res3, res4];
  if (stepStep.value === 1) {
    res = res1;
  } else if (stepStep.value === 2) {
    res = res2;
  } else if (stepStep.value === 3) {
    res = res3;
  } else if (stepStep.value === 4) {
    res = res4;
  }

  const resAnswer = await labStore.sendAlternativeDiffMatrices({
    matrices: res,
    step: stepStep.value,
  });

  if (resAnswer.data) {
    // answer.value = resAnswer.data;
    if (stepStep.value === 1) {
      answer1.value = resAnswer.data;
    } else if (stepStep.value === 2) {
      answer2.value = resAnswer.data;
    } else if (stepStep.value === 3) {
      answer3.value = resAnswer.data;
    } else if (stepStep.value === 4) {
      answer4.value = resAnswer.data;
      await labStore.incrementSecondStep();
    }
    stepStep.value += 1;
    localStorage.setItem(
      `answer-2-${stepStep.value}`,
      JSON.stringify(resAnswer.data)
    );
    labStore.getInfo();
  }
};

const columnsAnswer = ref([]);
const rows1Answer = ref([]);
const rows2Answer = ref([]);
const rows3Answer = ref([]);
const rows4Answer = ref([]);

const answer = computed(() => [
  answer1.value,
  answer2.value,
  answer3.value,
  answer4.value,
]);

watch(
  answer1,
  () => {
    localStorage.setItem('answer-2-1', JSON.stringify(answer1.value));
  },
  { deep: true }
);
watch(
  answer2,
  () => {
    localStorage.setItem('answer-2-2', JSON.stringify(answer2.value));
  },
  { deep: true }
);
watch(
  answer3,
  () => {
    localStorage.setItem('answer-2-3', JSON.stringify(answer3.value));
  },
  { deep: true }
);
watch(
  answer4,
  () => {
    localStorage.setItem('answer-2-4', JSON.stringify(answer4.value));
  },
  { deep: true }
);

watch(answer, () => {
  if (answer.value[0] !== null) {
    columnsAnswer.value = [];
    columnsAnswer.value.push({
      name: 'name',
      required: true,
      label: 'Название',
    });

    const answerColumnsLenght = answer.value[0]?.result[0].length;
    if (answerColumnsLenght) {
      for (let i = 0; i < answerColumnsLenght; i++) {
        columnsAnswer.value.push({
          name: i,
          required: true,
          label: `A${i + 1}`,
          style: 'width: 300px;',
        });
      }
    }

    const answer1 = answer.value[0]?.result;
    if (answer1) {
      rows1Answer.value = answer1.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer2 = answer.value[1]?.result;
    if (answer2) {
      console.log(answer2);
      rows2Answer.value = answer2?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer3 = answer.value[2]?.result;
    if (answer3) {
      rows3Answer.value = answer3?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer4 = answer.value[3]?.result;

    if (answer4) {
      rows4Answer.value = answer4?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }
  }
});

onMounted(() => {
  if (answer.value[0] !== null) {
    columnsAnswer.value = [];

    columnsAnswer.value.push({
      name: 'name',
      required: true,
      label: 'Название',
    });

    const answerColumnsLenght = answer.value[0]?.result[0].length;
    for (let i = 0; i < answerColumnsLenght; i++) {
      columnsAnswer.value.push({
        name: i,
        required: true,
        label: `A${i + 1}`,
        style: 'width: 300px;',
      });
    }

    const answer1 = answer.value[0]?.result;
    if (answer1) {
      stepStep.value = 2;
      rows1Answer.value = answer1.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer2 = answer.value[1]?.result;
    if (answer2) {
      stepStep.value = 3;
      rows2Answer.value = answer2?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer3 = answer.value[2]?.result;
    if (answer3) {
      stepStep.value = 4;
      rows3Answer.value = answer3?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }

    const answer4 = answer.value[3]?.result;

    if (answer4) {
      stepStep.value = 5;
      rows4Answer.value = answer4?.map((row, i) => {
        const resRow = [];
        for (let a in row) {
          if (a !== 'name') {
            resRow.push(row[a] === -1 ? '' : row[a]);
          }
        }
        resRow.name = { data: `A${i + 1}` };
        return resRow;
      });
    }
  }
});
</script>

<style lang="scss" scoped></style>
