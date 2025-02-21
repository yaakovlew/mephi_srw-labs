<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary flex justify-between items-center">
      <div>Определите оценку важности критериев</div>
      <q-toggle
        v-model="addInfoToggle"
        color="yellow"
        label="Доп. информация"
        icon="lightbulb"
      />
    </div>
    <div
      v-if="addInfoToggle"
      class="text-h6 text-primary flex column g-m justify-between"
    >
      <table>
        <tbody>
          <tr>
            <th class="row-info" scope="row">Важный</th>
            <td class="row-info">В</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">Очень важный</th>
            <td class="row-info">ОВ</td>
          </tr>
          <tr>
            <th class="row-info" scope="row">Не очень важный</th>
            <td class="row-info">НОВ</td>
          </tr>
        </tbody>
      </table>
      <div class="text-h5 text-primary" style="font-weight: 700">
        Использовать данные сокращенные варианты
      </div>
    </div>
    <q-table
      flat
      bordered
      title="Ответ"
      :rows="rows"
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
              v-if="col.name !== 'name' && isDone === false"
              v-model="props.row[col.name]"
              buttons
              v-slot="scope"
            >
              <q-input v-model="scope.value" />
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
    <q-table
      flat
      bordered
      title="Правильный ответ"
      :rows="rowsAnswer"
      :columns="columnsAnswer"
      row-key="name"
      binary-state-sort
      hide-bottom
      :separator="'cell'"
      v-else
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
            :class="{ 'error-field': props.row[col.name].flag === false }"
          >
            {{
              props.row[col.name] === -1
                ? ''
                : props.row[col.name].data
                ? String(props.row[col.name].data).toLocaleUpperCase()
                : props.row[col.name]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div v-if="answer">
      Количество полученных баллов за этап: {{ answer.percentage }} из
      {{ answer.max_mark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Ref, computed, onMounted, onUnmounted, ref, watch } from 'vue';
import { useLab3cStore } from 'src/stores/lab3c';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLab3cStore();

const isDone = computed(() => labStore.info?.step !== 1);

let columns = ref([]);

const rows = ref([]);

const variants = ['В', 'С', 'Н', 'ОН'];

const addInfoToggle = ref(false);

onMounted(() => {
  if (!props.variant) return;
  columns.value = [
    {
      name: 'weight',
      required: true,
      label: `Важность`,
      align: 'left',
    },
  ];

  columns.value.unshift({
    name: 'name',
    required: true,
    label: 'Название',
    align: 'left',
  });

  const userAnswers = localStorage.getItem('user-answers-2');
  if (userAnswers) {
    rows.value = JSON.parse(userAnswers);
  } else {
    props.variant?.criteria.forEach((criteria, i) => {
      const row = {};

      columns.value.forEach((column) => {
        const name = column.name;
        row.name = `С${i + 1}`;
        row[name] = variants[3];
      });

      rows.value.push(row);
    });
  }

  if (labStore.info?.step === 1) {
    localStorage.removeItem('answer-2');
    answer.value = null;
  }
});

watch(
  rows,
  () => {
    localStorage.setItem('user-answers-2', JSON.stringify(rows.value));
  },
  {
    deep: true,
  }
);

const isSend = ref(false);
const answer: Ref<Lab.AlternativeSetAnswer | null> = ref(
  localStorage.getItem('answer-2')
    ? JSON.parse(localStorage.getItem('answer-2')!)
    : null
);

const sendResult = async () => {
  const res = rows.value.map((row) => {
    const resRow = [];
    for (let a in row) {
      if (a !== 'name') {
        resRow.push(row[a].toString().toLowerCase());
      }
    }
    return resRow;
  });

  const finalRes = [];
  res.forEach((row) => {
    finalRes.push(row[0]);
  });

  const resAnswer = await labStore.sendCriteriaMatrix({ set: finalRes });
  if (resAnswer.data) {
    answer.value = resAnswer.data;
    localStorage.setItem('answer-2', JSON.stringify(answer.value));
    labStore.getInfo();
    labStore.normalStep = 3;
  }
};

const columnsAnswer = ref([]);
const rowsAnswer = ref([]);

watch(answer, () => {
  if (answer.value) {
    columnsAnswer.value.push({
      name: 'name',
      required: true,
      label: 'Название',
      align: 'left',
    });

    columnsAnswer.value.push({
      name: 'wei',
      required: true,
      label: `Важность`,
      align: 'left',
    });

    const values = answer.value?.result;

    values.forEach((value, index) => {
      const row = {
        wei: value,
        name: `C${index + 1}`,
      };
      rowsAnswer.value.push(row);
    });
  }
});

onMounted(() => {
  if (!answer.value) return;
  columnsAnswer.value.push({
    name: 'name',
    required: true,
    label: 'Название',
    align: 'left',
  });

  columnsAnswer.value.push({
    name: 'wei',
    required: true,
    label: `Важность`,
    align: 'left',
  });

  const values = answer.value?.result;

  values.forEach((value, index) => {
    const row = {
      wei: value,
      name: `C${index + 1}`,
    };
    rowsAnswer.value.push(row);
  });
});
</script>

<style lang="scss" scoped>
.row-info {
  width: 50%;
}

.row-info {
  border: 1px solid var(--q-primary);
  padding: 8px;
  text-align: left;
  text-transform: capitalize;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
  border: 2px solid var(--q-primary);
}
</style>
