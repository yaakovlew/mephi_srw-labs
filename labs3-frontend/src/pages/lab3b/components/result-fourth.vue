<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary">Постройте функциональное решение</div>
    <q-table
      v-for="(answer, index) in answers"
      :key="index"
      flat
      bordered
      title="Ответ"
      :rows="answer"
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
      v-if="!answer"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <template v-else>
      <q-table
        v-for="(answer, index) in finalAnswers"
        :key="index"
        flat
        bordered
        title="Правильный ответ"
        :rows="answer"
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
      <div>
        Количество полученных баллов за этап: {{ answer.percentage }} из
        {{ answer.max_mark }}
      </div>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useLab3bStore } from 'src/stores/lab3b';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  variant: Lab.Variant;
}>();

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

const isDone = computed(() => labStore.info?.step !== 3);

onMounted(() => {
  if (!props.variant) return;

  const userAnswers1 = localStorage.getItem('user-answers-4');
  if (userAnswers1) {
    answers.value = JSON.parse(userAnswers1);
  } else {
    const answer = [];
    props.variant.alternative.forEach((_, index) => {
      answer.push({
        ...JSON.parse(JSON.stringify(rowTemp.value)),
        name: `a${index + 1}`,
      });
    });
    answers.value.push(answer);
  }

  if (labStore.info?.step === 3) {
    localStorage.removeItem('answer-4');
    answer.value = null;
  }
});

watch(
  answers,
  () => {
    localStorage.setItem('user-answers-4', JSON.stringify(answers.value));
  },
  {
    deep: true,
  }
);

const answer: Ref<Lab3b.AllMatricesResult | null> = ref(
  localStorage.getItem('answer-4')
    ? JSON.parse(localStorage.getItem('answer-4')!)
    : null
);

const sendResult = async () => {
  const firstRow = answers.value[0];
  const res = firstRow.map((a) => {
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

  const resAnswer = await labStore.sendAllMatricesIntersection({
    matrix: res,
  });

  if (resAnswer.data) {
    answer.value = resAnswer.data;
    localStorage.setItem('answer-4', JSON.stringify(answer.value));
    labStore.getInfo();
  }
};

watch(answer, () => {
  if (answer.value) {
    const res = answer.value.result.map((a, index) => {
      const resO = { name: { data: `a${index + 1}` } };
      a.forEach((o) => {
        resO[`${o.X}`] = { data: o.Y, flag: o.Flag };
      });
      return resO;
    });
    finalAnswers.value = [res];
  }
});

onMounted(() => {
  if (!answer.value) return;
  const res = answer.value.result.map((a, index) => {
    const resO = { name: { data: `a${index + 1}` } };
    a.forEach((o) => {
      resO[`${o.X}`] = { data: o.Y, flag: o.Flag };
    });
    return resO;
  });
  finalAnswers.value = [res];
});
</script>

<style lang="scss" scoped></style>
