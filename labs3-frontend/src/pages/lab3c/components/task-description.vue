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
      <q-table
        style="margin-top: 20px; margin-bottom: 20px"
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
      <div class="text-h6 color-primary">Важность критериев</div>
      <div class="flex column g-xl">
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
            <Chart
              class="hc"
              :options="
                createOptionFromPoint(criteria.points, criteria.importance)
              "
              ref="chart"
              style="width: 100%"
            ></Chart>
          </div>
        </div>
        <div class="text-h6 color-primary">Важность альтернатив</div>
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
            <Chart
              class="hc"
              :options="
                createOptionFromPoint(criteria.points, criteria.importance)
              "
              ref="chart"
              style="width: 100%"
            ></Chart>
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
import { Chart } from 'highcharts-vue';

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

const capitalizeFirstLetter = (str: string) =>
  str.charAt(0).toUpperCase() + str.slice(1);

const createOptionFromPoint = (points: Lab.Point[], name: string) => ({
  series: [
    {
      data: points.map((p) => [p.X, p.Y]),
      name: '',
    },
  ],
  tooltip: {
    headerFormat: '',
    pointFormat: '{point.x}; {point.y}',
  },
  xAxis: {
    tickInterval: 0.01,
    reversed: false,
    maxPadding: 0.05,
    showLastLabel: true,
    accessibility: {
      rangeDescription: 'Range: 0 to 1',
    },
    max: 1.05,
    min: 0,
  },
  title: {
    text: capitalizeFirstLetter(name),
  },
});


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
