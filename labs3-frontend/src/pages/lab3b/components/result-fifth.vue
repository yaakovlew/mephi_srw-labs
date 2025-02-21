<template>
  <div class="q-pa-md flex column g-xl">
    <div class="text-h6 text-primary">
      Определите итоговую оценку альтернатив, выберите наиболее оптимальную
      альтернативу
    </div>
    <div
      v-for="(answerTable, indexTable) in answers"
      :key="indexTable"
      class="flex column g-lg"
    >
      <div class="text-h6 text-primary">
        D<span style="font-size: 12px; margin-top: 3px">
          {{ indexTable + 1 }}
        </span>
      </div>
      <div
        v-for="(a, index) in answerTable"
        :key="index"
        class="flex g-m column"
      >
        <q-btn
          v-if="!answer"
          @click="deleteTable(indexTable, index)"
          label="Удалить таблицу"
          color="negative"
          icon="delete"
          flat
          style="max-width: 300px"
          class="q-mt-md"
        />
        <div class="flex g-m row">
          <span class="text-h6 text-primary">
            E<span style="font-size: 12px; margin-top: 3px">
              {{ index + 1 }}α
            </span>
            :
          </span>
          <span class="text-h6 text-primary" v-html="specialSymbol" />
          <div v-for="(element, index) in a.set" :key="index" class="flex g-m">
            <div class="text-h6 text-primary">
              {{ element }}
              <q-popup-edit
                v-model.number="a.set[index]"
                auto-save
                v-slot="scope"
              >
                <q-input
                  v-model.number="scope.value"
                  @blur="scope.value = roundToNDigits(scope.value)"
                  type="number"
                  dense
                  autofocus
                  counter
                  @keyup.enter="scope.set"
                />
              </q-popup-edit>
            </div>
            <q-separator
              vertical
              style="height: 100%"
              v-if="index !== a.set.length - 1"
            />
          </div>
          <span class="text-h6 text-primary" v-html="specialSymbolClosed" />
        </div>
        <div>
          <q-btn
            v-if="!answer"
            @click="addToSet(indexTable, index)"
            label="Добавить элемент"
            color="primary"
            icon="add"
            flat
          />
          <q-btn
            v-if="!answer"
            @click="removeFromSet(indexTable, index)"
            label="Удалить последний элемент"
            color="negative"
            icon="delete"
            flat
          />
        </div>
        <q-input
          v-model.number="a.powerful"
          label="Мощность"
          flat
          filled
          :disabled="answer"
        />
        <q-input
          v-model.number="a.delta"
          label="Дельта"
          flat
          filled
          :disabled="answer"
        />
        <q-separator v-if="index !== answers.length - 1" color="primary" />
      </div>
      <div v-if="isDone === false">
        <q-btn
          label="Добавить таблицу"
          color="primary"
          @click="createTable(indexTable)"
        />
      </div>
    </div>
    <q-btn
      v-if="!answer"
      label="Отправить"
      color="primary"
      @click="sendResult"
    />
    <div
      v-for="(answersTable, indexTable) in columnsAnswer"
      :key="indexTable"
      class="flex column g-xl"
    >
      <div class="text-h6 text-primary">D{{ indexTable + 1 }}</div>
      <div
        v-for="(a, index) in answersTable"
        :key="index"
        class="flex g-m column"
      >
        <div class="flex g-m row">
          <span class="text-h6 text-primary" v-html="specialSymbol" />
          <div v-for="(element, index) in a.set" :key="index" class="flex g-m">
            <div class="text-h6 text-primary">
              {{ element }}
            </div>
            <q-separator
              vertical
              style="height: 100%"
              v-if="index !== a.set.length - 1"
            />
          </div>
          <span class="text-h6 text-primary" v-html="specialSymbolClosed" />
        </div>
        <div class="text-h6 text-primary">Мощность: {{ a.powerful }}</div>
        <div class="text-h6 text-primary">Дельта: {{ a.delta }}</div>
        <q-separator v-if="index !== answers.length - 1" color="primary" />
      </div>
    </div>
    <div v-if="answer">
      Количество полученных баллов за этап: {{ answer.percentage }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useLab3bStore } from '../../../stores/lab3b';

const props = defineProps<{
  variant: Lab.Variant;
}>();

const labStore = useLab3bStore();

const isDone = computed(() => labStore.info?.step !== 4);

let columns = ref([]);

const answers: Ref<Lab3b.LevelSet[][]> = ref([]);

const specialSymbol = '{';
const specialSymbolClosed = '}';

const columnsToCreateTable = ref('');
const columnsCount = ref(0);

const createTable = (index: number) => {
  answers.value[index].push({
    set: [0],
    delta: 0,
    powerful: 0,
  });
};

const deleteTable = (indexTable: number, index: number) => {
  answers.value[indexTable].splice(index, 1);
};

const addToSet = (indexTable: number, index: number) => {
  answers.value[indexTable][index].set.push(0);
};

const removeFromSet = (indexTable: number, index: number) => {
  answers.value[indexTable][index].set.pop();
};

watch(
  answers,
  () => {
    localStorage.setItem('user-answers-5', JSON.stringify(answers.value));
  },
  {
    deep: true,
  }
);

const rows = ref([]);

onMounted(() => {
  if (!props.variant) return;

  const userAnswers = localStorage.getItem('user-answers-5');
  if (userAnswers) {
    answers.value = JSON.parse(userAnswers);
  } else {
    props.variant.alternative.forEach(() => {
      answers.value.push([]);
    });
  }

  if (labStore.info?.step === 4) {
    localStorage.removeItem('answer-5');
    answer.value = null;
  }
});

const isSend = ref(false);
const answer: Ref<Lab3b.SendLevelSetResult | null> = ref(
  localStorage.getItem('answer-5')
    ? JSON.parse(localStorage.getItem('answer-5')!)
    : null
);

const sendResult = async () => {
  const res = answers.value;

  const resAnswer = await labStore.sendLevelSet({ answer_level_set: res });
  if (resAnswer.data) {
    answer.value = resAnswer.data;
    localStorage.setItem('answer-5', JSON.stringify(answer.value));
    labStore.getInfo();
  }
};

const columnsAnswer: Ref<Lab3b.LevelSet[][]> = ref([]);

watch(answer, () => {
  if (answer.value) {
    columnsAnswer.value = answer.value.result;
  }
});

onMounted(() => {
  if (!answer.value) return;
  columnsAnswer.value = answer.value.result;
});
</script>

<style lang="scss" scoped></style>
