<template>
  <div class="flex column g-m">
    <div class="text-h5 text-primary">Ответ</div>
    <q-table
      flat
      bordered
      title="Ответ"
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
              (model[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                props.pageIndex
              ]
            }}
            <q-popup-edit
              v-if="col.name !== 'name' && !isDone"
              v-model="(model[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                props.pageIndex
              ]"
              buttons
              v-slot="scope"
            >
              <q-input
                v-model.number="(model[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                props.pageIndex
              ]"
                @blur="
                  (model[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                    props.pageIndex
                  ] = roundToNDigits(
                    (model[col.name as keyof Lab1a.MatrixAnswer] as number[])[
                      props.pageIndex
                    ]
                  )
                "
                type="number"
                dense
                autofocus
                @keydown.enter="scope.set"
              />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <div>
      <q-input
        type="number"
        v-model.number="model.eigenvalue"
        label="λmax"
        :disable="isDone"
      />
      <q-input
        type="number"
        v-model.number="model.consistencyIndex"
        label="ИС"
        :disable="isDone"
      />
      <q-input
        type="number"
        v-model.number="model.consistencyRatio"
        label="ОС"
        :disable="isDone"
      />
    </div>
    <q-btn
      v-if="!isDone"
      flat
      color="primary"
      label="Отправить"
      @click="$emit('confirm')"
    />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue';
import { roundToNDigits } from 'src/utils/round-to-n-digits';
import { Lab1a } from 'src/models/lab/lab1a';

const props = defineProps<{
  modelValue: Lab1a.MatrixAnswer;
  isDone: boolean;
}>();

const model = ref<Lab1a.MatrixAnswer>(props.modelValue);

const emit = defineEmits<{
  (e: 'update:modelValue', value: Lab1a.MatrixAnswer): void;
  (e: 'confirm'): void;
}>();

watch(
  model,
  () => {
    emit('update:modelValue', model.value);
  },
  { deep: true }
);

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

  rows.value = model.value.weightVector.map(() => ({}));
});
</script>

<style lang="scss" scoped></style>
