<template>
  <div class="flex column g-m">
    <q-table
      flat
      bordered
      :rows="rowsWeights"
      :columns="(columnsWeights as any)"
      row-key="name"
      hide-bottom
      :separator="'cell'"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td v-for="col in props.cols" :key="col.name" :props="props">
            {{ props.row[col.name] }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
    <q-table
      title="Ответ"
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
                ? model.alternatives[props.pageIndex]
                : props.row[col.name]
            }}
            <q-popup-edit
              v-if="col.name !== 'name' && !isDone"
              v-model="model.alternatives[props.pageIndex]"
              buttons
              v-slot="scope"
            >
              <q-input
                v-model.number="model.alternatives[props.pageIndex]"
                @blur="
                  model.alternatives[props.pageIndex] = roundToNDigits(
                    model.alternatives[props.pageIndex]
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
    <q-select
      v-model="choosenAlternative"
      :options="alternativesOptions"
      label="Выбранная альтернатива"
      :disable="isDone"
    />
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
import { Lab1a } from 'src/models/lab/lab1a';
import { computed, onMounted, ref, watch } from 'vue';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

const props = defineProps<{
  alternatives: string[];
  modelValue: Lab1a.PrioritiesAlternatives;
  isDone: boolean;
  criterias: string[];
  weights: number[][];
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: Lab1a.PrioritiesAlternatives): void;
  (e: 'confirm'): void;
}>();

const model = ref<Lab1a.PrioritiesAlternatives>(props.modelValue);

const alternativesOptions = computed(() =>
  props.alternatives.map((a, i) => ({ label: a, value: i }))
);

const choosenAlternative = ref(
  model.value.choosenAlternativeIndex !== null
    ? alternativesOptions.value[model.value.choosenAlternativeIndex]
    : alternativesOptions.value[0]
);

watch(
  model,
  () => {
    emit('update:modelValue', model.value);
  },
  { deep: true }
);

watch(
  choosenAlternative,
  () => {
    model.value.choosenAlternativeIndex = choosenAlternative.value.value;
  },
  { deep: true }
);

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

const weightsTitle = 'Веса критериев';

const columnsWeights = [
  {
    name: 'name',
    required: true,
    field: 'name',
    style: nameStyle,
  },
  {
    name: 'weights',
    required: true,
    field: 'weights',
    label: weightsTitle,
  },
];

const rowsWeights = ref<Array<Record<string, unknown>>>([]);

const rowsAlternatives = ref<Array<Record<string, unknown>>>([]);

onMounted(() => {
  props.alternatives.forEach((alternative, i) => {
    rowsAlternatives.value.push({
      name: alternative,
      priority: 0,
    });
  });

  props.criterias.forEach((criteria, i) => {
    rowsWeights.value.push({
      name: criteria,
      weights: props.weights[i][0],
    });
  });
});
</script>

<style lang="scss" scoped></style>
