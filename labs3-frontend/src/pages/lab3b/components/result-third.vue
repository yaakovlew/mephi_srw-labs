<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary flex justify-between items-center">
      Постройте нечеткие отношения с использованием правила Лукасевича
      <q-toggle
        v-model="addInfoToggle"
        color="yellow"
        label="Правило Лукасевича"
        icon="lightbulb"
      />
    </div>
    <div v-if="addInfoToggle" class="flex items-center text-h6 text-primary">
      Правило Лукасевича: <img :src="Pravilo" />
    </div>
    <template v-for="(answer1, index) in answers" :key="index">
      <q-table
        v-if="index < currentStep"
        flat
        bordered
        title="Ответ"
        :rows="answer1"
        :columns="columnsTemp"
        row-key="name"
        binary-state-sort
        hide-bottom
        :separator="'cell'"
      >
        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td v-for="col in props.cols" :key="col.name" :props="props">
              {{ props.row[col.name] === -1 ? '' : props.row[col.name] }}
            </q-td>
          </q-tr>
        </template>
      </q-table>
      <q-table
        v-if="index < currentStep"
        flat
        bordered
        title="Правильный ответ"
        :rows="finalAnswers[index]"
        :columns="columnsTemp"
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
      <div v-if="index < currentStep">
        Количество полученных баллов за шаг: {{ answer[index].percentage }} из
        {{ answer[index].max_mark }}
      </div>
    </template>
    <q-table
      v-if="currentStep < 3"
      flat
      bordered
      title="Ответ"
      :rows="answers[currentStep]"
      :columns="columnsTemp"
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
      v-if="currentStep < 3"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <div v-if="answer" style="font-size: 18px">
      Общее количество полученных баллов за шаги: {{ totalSum.sum }} из
      {{ totalSum.maxCount }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useLab3bStore } from 'src/stores/lab3b';
import { roundToNDigits } from 'src/utils/round-to-n-digits';
import Pravilo from '../pravilo.png';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const addInfoToggle = ref(false);

const columnsTemp = ref([
  {
    name: 'name',
    field: 'name',
    label: '',
  },
  {
    name: '0',
    field: '0',
    label: '0',
    style: 'width: 100px',
  },
  {
    name: '0.1',
    field: '0.1',
    label: '0.1',
    style: 'width: 100px',
  },
  {
    name: '0.2',
    field: '0.2',
    label: '0.2',
    style: 'width: 100px',
  },
  {
    name: '0.3',
    field: '0.3',
    label: '0.3',
    style: 'width: 100px',
  },
  {
    name: '0.4',
    field: '0.4',
    label: '0.4',
    style: 'width: 100px',
  },
  {
    name: '0.5',
    field: '0.5',
    label: '0.5',
    style: 'width: 100px',
  },
  {
    name: '0.6',
    field: '0.6',
    label: '0.6',
    style: 'width: 100px',
  },
  {
    name: '0.7',
    field: '0.7',
    label: '0.7',
    style: 'width: 100px',
  },
  {
    name: '0.8',
    field: '0.8',
    label: '0.8',
    style: 'width: 100px',
  },
  {
    name: '0.9',
    field: '0.9',
    label: '0.9',
    style: 'width: 100px',
  },
  {
    name: '1',
    field: '1',
    label: '1',
    style: 'width: 100px',
  },
]);

const rowTemp = ref({
  name: 'Y',
  '0': 0,
  '0.1': 0,
  '0.2': 0,
  '0.3': 0,
  '0.4': 0,
  '0.5': 0,
  '0.6': 0,
  '0.7': 0,
  '0.8': 0,
  '0.9': 0,
  '1': 0,
});

const answers = ref([]);

const finalAnswers = ref([]);

const labStore = useLab3bStore();

const isDone = computed(() => labStore.info?.step !== 2);

const currentStep = ref(0);

onMounted(() => {
  if (!props.variant) return;

  const userAnswers1 = localStorage.getItem('user-answers-3');
  if (userAnswers1) {
    answers.value = JSON.parse(userAnswers1);
  } else {
    props.variant.rule.forEach((_, index) => {
      const answer = [];
      props.variant.alternative.forEach((_, i) => {
        answer.push({
          ...JSON.parse(JSON.stringify(rowTemp.value)),
          name: `D${index + 1}(a${i + 1})`,
        });
      });
      answers.value.push(answer);
    });
  }

  if (localStorage.getItem('step-3')) {
    currentStep.value = parseInt(localStorage.getItem('step-3')) ?? 0;
  }

  if (currentStep.value === 0) {
    localStorage.removeItem('answer-3');
    answer.value = [];
  }
});

watch(
  answers,
  () => {
    localStorage.setItem('user-answers-3', JSON.stringify(answers.value));
  },
  {
    deep: true,
  }
);

const isSend = ref(false);
const answer: Ref<Lab3b.AllMatricesResult[] | null> = ref(
  localStorage.getItem('answer-3')
    ? JSON.parse(localStorage.getItem('answer-3')!)
    : []
);

const sendResult = async () => {
  const res = answers.value[currentStep.value].map((a) => {
    const resRow = [];
    Object.keys(a).forEach((key) => {
      if (key !== 'name') {
        const o = {
          X: roundToNDigits(key),
          Y: roundToNDigits(a[key]),
        };
        resRow.push(o);
      }
    });
    return resRow;
  });

  const resAnswer = await labStore.sendAllMatrices({
    matrices: res,
    step: currentStep.value + 1,
  });

  if (resAnswer.data) {
    currentStep.value = currentStep.value + 1;
    answer.value?.push(resAnswer.data);
    localStorage.setItem('answer-3', JSON.stringify(answer.value));
    if (currentStep.value === answers.value.length) {
      await labStore.increment2Step();
    }
    await labStore.getInfo();
  }
};

watch(
  currentStep,
  () => {
    localStorage.setItem('step-3', JSON.stringify(currentStep.value));
  },
  {
    deep: true,
  }
);

const totalSum = computed(() => {
  let sum = 0;
  let maxCount = 0;
  answer.value?.forEach((item) => {
    if (item) {
      sum += item.percentage;
      maxCount += item.max_mark;
    }
  });

  localStorage.setItem(
    'answer-3-sum',
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
  answer,
  () => {
    if (!answer.value) return;
    finalAnswers.value = answer.value?.map((ans, index) => {
      const res = ans.result.map((a, i) => {
        const resO = { name: { data: `D${index + 1}(a${i + 1})` } };
        a.forEach((o) => {
          console.log(o);
          resO[`${o.X}`] = { data: o.Y, flag: o.Flag };
        });
        console.log(resO);
        return resO;
      });
      return res;
    });
  },
  {
    deep: true,
  }
);

onMounted(() => {
  if (!answer.value) return;
  finalAnswers.value = answer.value?.map((ans, index) => {
    const res = ans.result.map((a, i) => {
      const resO = { name: { data: `D${index + 1}(a${i + 1})` } };
      a.forEach((o) => {
        console.log(o);
        resO[`${o.X}`] = { data: o.Y, flag: o.Flag };
      });
      console.log(resO);
      return resO;
    });
    return res;
  });
});
</script>

<style lang="scss" scoped></style>
