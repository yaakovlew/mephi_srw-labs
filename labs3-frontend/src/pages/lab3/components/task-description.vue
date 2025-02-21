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
      <div v-if="currentMode === 3" class="flex column g-xl">
        <table>
          <thead>
            <tr>
              <th scope="col"></th>
              <th scope="col"></th>
              <th scope="col">0</th>
              <th scope="col">1</th>
              <th scope="col">0</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="criteria in variant?.importance_criteria"
              :key="criteria.importance"
            >
              <th scope="row">{{ criteria.importance }}</th>
              <td>{{ criteria.short_importance_name.toUpperCase() }}</td>
              <td v-for="value in criteria.points" :key="value.X">
                {{ value.X }}
              </td>
            </tr>
          </tbody>
        </table>
        <div class="flex row g-m items-center justify-between">
          <div
            v-for="criteria in variant?.importance_criteria"
            :key="criteria.importance"
          >
            {{ criteria.short_importance_name.toLocaleUpperCase() }}
            <div :id="`chart-${criteria.importance.replace(/\s/g, '')}`"></div>
            <!-- <Chart :data="(criteria.points as any[])" :margin="margin">
              <template #layers>
                <Grid strokeDasharray="1,1" />
                <Line :dataKeys="['X', 'Y']" />
              </template>
            </Chart> -->
          </div>
        </div>
        <table>
          <thead>
            <tr>
              <th scope="col"></th>
              <th scope="col"></th>
              <th scope="col">0</th>
              <th scope="col">1</th>
              <th scope="col">0</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="criteria in variant?.importance_alternative"
              :key="criteria.importance"
            >
              <th scope="row">{{ criteria.importance }}</th>
              <td>{{ criteria.short_importance_name.toUpperCase() }}</td>
              <td v-for="value in criteria.points" :key="value.X">
                {{ value.X }}
              </td>
            </tr>
          </tbody>
        </table>
        <div class="flex row g-m items-center justify-between">
          <div
            v-for="criteria in variant?.importance_alternative"
            :key="criteria.importance"
          >
            {{ criteria.short_importance_name.toLocaleUpperCase() }}
            <Chart :data="(criteria.points as any[])" :margin="margin">
              <template #layers>
                <Grid strokeDasharray="1,1" />
                <Line :dataKeys="['X', 'Y']" />
              </template>
            </Chart>
          </div>
        </div>
      </div>
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
}>();

const varaintNumber = computed(() => props.variant?.number);

function getImageUrl(name: number) {
  return new URL(`../${varaintNumber.value}/${name}.png`, import.meta.url).href;
}
function getImageFuncUrl(name: number) {
  return new URL(`../${varaintNumber.value}/funcs/${name}.png`, import.meta.url)
    .href;
}

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

const currentMode = ref(modes[0]);

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
