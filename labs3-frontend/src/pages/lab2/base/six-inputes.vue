<template>
  <div class="flex column g-m">
    <q-table
      flat
      bordered
      :title="title"
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
            <template v-if="model">
              {{ model ? (model[col.name] as number) : '' }}
              <q-popup-edit
                v-if="col.name !== 'name' && !isDone"
                v-model.number="model[col.name]"
                buttons
                v-slot="scope"
              >
                <q-input
                  v-model.number="model[col.name]"
                  @blur="model[col.name] = roundToNDigits(model[col.name])"
                  type="number"
                  dense
                  autofocus
                  @keydown.enter="scope.set"
                />
              </q-popup-edit>
            </template>
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  modelValue: number[];
  isDone: boolean;
  title: string;
}>();

const model = ref<number[]>(props.modelValue);

const columns = ref<Array<Record<string, unknown>>>([]);
const rows = ref<Array<Record<string, unknown>>>([{}]);

onMounted(() => {
  const style = `width: ${100 / 6}%`;
  columns.value = [
    {
      name: '0',
      label: 'Критерий 1',
      field: '0',
      style,
    },
    {
      name: '1',
      label: 'Критерий 2',
      field: '1',
      style,
    },
    {
      name: '2',
      label: 'Критерий 3',
      field: '2',
      style,
    },
    {
      name: '3',
      label: 'Критерий 4',
      field: '3',
      style,
    },
    {
      name: '4',
      label: 'Критерий 5',
      field: '4',
      style,
    },
    {
      name: '5',
      label: 'Критерий 6',
      field: '5',
      style,
    },
  ];
});
</script>

<style lang="scss" scoped></style>
