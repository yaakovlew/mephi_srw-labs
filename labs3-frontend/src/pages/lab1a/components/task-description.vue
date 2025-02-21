<template>
  <div>
    <banner-component class="text-primary text">
      <template #title>
        <div>Ваиант: {{ variantNumber }}</div>
      </template>
      <div>
        <span class="text-main"> Цель: </span> выбор экономической стратегии
      </div>
      <div class="text-main">Критерии:</div>
      <ul class="q-py-none q-my-none">
        <li>
          <div>Производительность (П)</div>
        </li>
        <li>
          <div>Конкурентноспособность (К)</div>
        </li>
        <li>
          <div>Осуществимость стратегии (О)</div>
        </li>
        <li>
          <div>Риск (Р)</div>
        </li>
      </ul>
      <div class="text-main">Альтернативы:</div>
      <ul class="q-py-none q-my-none">
        <li>
          <div>Стратегия А</div>
        </li>
        <li>
          <div>Стратегия Б</div>
        </li>
        <li>
          <div>Стратегия В</div>
        </li>
      </ul>
      <div class="text-main">По итогам опроса ЛПР составлены матрицы:</div>
      <div>
        <ol class="q-py-none q-my-none">
          <li>
            <div>По цели: размер 4*4, одна штука</div>
          </li>
          <li>
            <div>По каждому критерию: размер 3*3, четыре штуки</div>
          </li>
        </ol>
      </div>
      <div class="text-main">Задание:</div>

      <div>
        <ul class="q-py-none q-my-none">
          <li>
            <div>для каждой матрицы рассчитать веса критериев;</div>
          </li>
          <li>
            <div>для каждой матрицы рассчитать ИС и ОС;</div>
          </li>
          <li>
            <div>
              для несогласованных матриц (ОС менее 10%) необходимо произвести их
              пересогласование;
            </div>
          </li>
          <li>
            <div>
              рассчитать приоритеты альтернатив относительно их пригодности для
              достижения поставленной цели [kA, kБ, kВ].
            </div>
          </li>
        </ul>
      </div>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import { Lab1a } from 'src/models/lab/lab1a';

const props = defineProps<{
  variant: Lab1a.Variant | null;
  variantNumber?: number;
}>();
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
