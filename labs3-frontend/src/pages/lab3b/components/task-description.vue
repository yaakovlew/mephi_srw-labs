<template>
  <div>
    <banner-component class="text-primary text">
      <div><span class="text-main"> Задание: </span> {{ variant?.task }}</div>
      <div class="text-main">Критерии:</div>
      <ul class="q-py-none q-my-none">
        <li
          v-for="(criteria, index) in variant?.criteria"
          :key="criteria.definition"
        >
          <div>
            {{ criteria.definition }}
          </div>
          <div>Вес: {{ criteria.weight }}</div>
          <div>
            {{ criteria.extra_info }}
          </div>
          <img :src="getImageUrl(index + 1)" />
          <div
            v-if="criteria.func"
            style="display: flex; align-items: center"
            class="g-m"
          >
            Функция:
            <img :src="getImageFuncUrl(index + 1)" style="width: 100px" />
          </div>
          <!-- <div class="flex g-m">
            <div>Оценки:</div>
            <div>
              <div v-for="mark in criteria.func_mark" :key="mark.name">
                {{ mark.func }} - {{ mark.name }}
              </div>
            </div>
          </div> -->
        </li>
      </ul>
      <div class="text-main">Альтернативы:</div>
      <ul class="q-py-none q-my-none">
        <li
          v-for="alternative in variant?.alternative"
          :key="alternative.description"
        >
          <div>
            {{ alternative.description }}
          </div>
          <div class="flex g-m">
            <div>
              <div v-for="(constant, i) in constantsToMarks" :key="i">
                {{ constant }}:
                {{
                  i < 2
                    ? alternative.criteria_count[i].count
                    : `${alternative.criteria_count[i].count} ${alternative.criteria_count[i].value}`
                }}
              </div>
            </div>
          </div>
        </li>
      </ul>
      <div v-if="currentMode === 2">
        <div class="text-main">Правила:</div>
        <ul class="q-py-none q-my-none">
          <li v-for="rule in variant?.rule" :key="rule.name">
            {{ rule.name }}
          </li>
        </ul>
      </div>
      <q-table
        style="margin-top: 10px"
        flat
        bordered
        title="Оценки альтернатив по критериям"
        :rows="rows"
        :columns="columns"
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
              class="border-black"
              style="border: 1px solid rgba(0, 0, 0, 0.12)"
            >
              {{ props.row[col.name] === -1 ? '' : props.row[col.name] }}
            </q-td>
          </q-tr>
        </template>
      </q-table>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import BannerComponent from 'src/components/BannerComponent.vue';
import { computed, onMounted, ref } from 'vue';
import { Chart, Grid, Line } from 'vue3-charts';
import ApexCharts from 'apexcharts';

const props = defineProps<{
  variant?: Lab.Variant;
  matrix?: number[][];
}>();

const varaintNumber = computed(() => props.variant?.number);

function getImageUrl(name: number) {
  return new URL(`../${varaintNumber.value}/${name}.png`, import.meta.url).href;
}

function getImageFuncUrl(name: number) {
  return new URL(`../${varaintNumber.value}/funcs/${name}.png`, import.meta.url)
    .href;
}

const columns = computed(() => {
  const res = [{ name: 'name', label: '' }];
  if (props.matrix && props.matrix[0]) {
    const columnsCount = props.matrix[0].length;

    for (let i = 0; i < columnsCount; i++) {
      res.push({
        name: i,
        required: true,
        label: `A${i + 1}`,
      });
    }
  }
  return res;
});

const rows = computed(() => {
  let res = [];
  if (props.matrix) {
    res = props.matrix.map((row, i) => {
      const resRow = [];
      for (let a in row) {
        if (a !== 'name') {
          resRow.push(row[a] === -1 ? '' : row[a]);
        }
      }
      resRow.name = `С${i + 1}`;
      return resRow;
    });
  }
  return res;
});

const modes = [1, 2, 3];

var options = {
  series: [
    {
      name: 'Peter',
      data: [5, 5, 10, 8, 7, 5, 4, null, null, null, 10, 10, 7, 8, 6, 9],
    },
  ],
  chart: {
    height: 350,
    type: 'line',
    zoom: {
      enabled: false,
    },
    animations: {
      enabled: false,
    },
  },
  stroke: {
    width: [5, 5, 4],
    curve: 'straight',
  },
  labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16],
  title: {
    text: 'Missing data (null values)',
  },
  xaxis: {},
};

onMounted(() => {
  for (let chart of charts) {
    new ApexCharts(document.querySelector(chart.id), chart.options).render();
  }
});

type ChartOptions = {
  series: {
    name: string;
    data: (number | null)[];
  }[];
  chart: {
    height: number;
    type: string;
    zoom: {
      enabled: boolean;
    };
    animations: {
      enabled: boolean;
    };
  };
  stroke: {
    width: number[];
    curve: string;
  };
  labels: (number | string)[];
  title: {
    text: string;
  };
  xaxis: Record<string, unknown>;
};

type ChartData = {
  id: string;
  options: ChartOptions;
};

const charts: ChartData[] = [];

onMounted(() => {
  if (!props.variant) return;
  for (let criteria of props.variant?.importance_criteria) {
    const id = '#chart-' + criteria.importance.replace(/\s/g, '');
    const options: ChartOptions = {
      chart: {
        height: 350,
        type: 'line',
        zoom: {
          enabled: false,
        },
        animations: {
          enabled: false,
        },
      },
      stroke: {
        width: [5, 5, 4],
        curve: 'straight',
      },
      title: {
        text: criteria.short_importance_name,
      },
      xaxis: {},
      series: [
        {
          name: criteria.short_importance_name,
          data: criteria.points.map((p) => p.Y),
        },
      ],
      labels: criteria.points.map((p) => p.X).map((p) => p.toString()),
    };
    charts.push({
      id,
      options,
    });
  }
});

const margin = ref({
  left: 0,
  top: 20,
  right: 20,
  bottom: 0,
});

const currentMode = ref(modes[1]);

const constantsToMarks = [
  'Уникальность',
  'Конкрентноспособность',
  'Аудитория',
  'Стоимость',
];
</script>

<style lang="scss" scoped>
.text {
  font-size: 18px;
  font-weight: 500;
}

.text-main {
  font-weight: 700;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
  border: 2px solid var(--q-primary);
}

th {
  border: 1px solid var(--q-primary);
  padding: 8px;
  text-align: left;
  text-transform: capitalize;
}

td {
  border: 1px solid var(--q-primary);
  padding: 8px;
  text-align: left;
}

.alternative {
  display: flex;
  flex-direction: column;
}
</style>
