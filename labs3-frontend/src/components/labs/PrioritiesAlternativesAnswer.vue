<template>
  <div class="flex column g-m">
    <q-table
      title="Правильный ответ"
      flat
      bordered
      :rows="rowsAlternatives"
      :columns="(columnsAlternatievs as any)"
      row-key="name"
      hide-bottom
      :separator="'cell'"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-for="col in props.cols" :key="col.name" :props="props">
            {{
              col.name !== 'name'
                ? correctAnswer.resultPriorities[props.pageIndex]
                : props.row[col.name]
            }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div class="flex column g-m text-h6 text-primary">
      Альтернатива: {{ rightAlternative }}
    </div>
    <div class="flex column g-m text-h6 text-primary">
      Баллов за этап {{ mark }} из {{ maxMark }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab1a } from 'src/models/lab/lab1a';
import { computed, onMounted, ref } from 'vue';

const props = defineProps<{
  alternatives: string[];
  correctAnswer: Lab1a.VariantAnswer;
  mark: number;
  maxMark: number;
}>();

const prioritiesTitle = 'Приоритеты альтернатив';

const nameStyle = 'width: 200px';

const columnsAlternatievs = [
  {
    name: 'name',
    required: true,
    field: 'name',
    style: nameStyle,
  },
  {
    name: 'priority',
    required: true,
    field: 'priority',
    label: prioritiesTitle,
  },
];

const rightAlternative = computed(
  () => props.alternatives[props.correctAnswer.bestAlternative]
);

const rowsAlternatives = ref<Array<Record<string, unknown>>>([]);

onMounted(() => {
  props.alternatives.forEach((alternative, i) => {
    rowsAlternatives.value.push({
      name: alternative,
      priority: 0,
    });
  });
});
</script>

<style lang="scss" scoped></style>
