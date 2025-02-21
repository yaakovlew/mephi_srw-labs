<template>
  <div class="flex column g-m">
    <div class="text-h5 text-primary">Правильный ответ</div>
    <q-table
      flat
      bordered
      title="Правильный ответ"
      :rows="rows"
      :columns="(columns as any)"
      row-key="name"
      binary-state-sort
      hide-bottom
      :separator="'cell'"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-for="col in props.cols" :key="col.name" :props="props">
            {{
              (matrix[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                props.pageIndex
              ]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div class="flex column g-m text-h6 text-primary">
      <div>λmax: {{ matrix.eigenvalue }}</div>
      <div>ИС: {{ matrix.consistencyIndex }}</div>
      <div>ОС: {{ matrix.consistencyRatio }}</div>
    </div>
    <div class="flex column g-m text-h6 text-primary">
      Баллов за этап {{ mark }} из {{ maxMark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab1a } from 'src/models/lab/lab1a';
import { onMounted, ref } from 'vue';

const props = defineProps<{
  matrix: Lab1a.MatrixAnswer;
  mark: number;
  maxMark: number;
}>();

const columns = ref<Array<Record<string, unknown>>>([]);
const rows = ref<Array<Record<string, unknown>>>([]);

onMounted(() => {
  const style = 'width: 25%';
  columns.value = [
    {
      name: 'priorityVector',
      label: 'X',
      field: 'priorityVector',
      style,
    },
    {
      name: 'weightVector',
      label: 'w',
      field: 'weightVector',
      style,
    },
    {
      name: 'matrixWeightVector',
      label: 'M*w',
      field: 'matrixWeightVector',
      style,
    },
    {
      name: 'lambdaVector',
      label: 'λ*w',
      field: 'lambdaVector',
      style,
    },
  ];

  rows.value = props.matrix.weightVector.map(() => ({}));
});
</script>

<style lang="scss" scoped></style>
