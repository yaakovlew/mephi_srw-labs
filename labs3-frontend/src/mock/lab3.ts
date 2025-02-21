import { Lab } from 'src/models/lab/lab';

export const lab3Variant: Lab.LabVariant = {
  variant: {
    number: 1,
    task: 'выбор нового товара для продвижения.',
    criteria: [
      {
        definition:
          'С1 Уникальность – показатель количества аналогичных предложений.',
        extra_info:
          '(градаций 5: очень низкий, низкий, средний, высокий, очень высокий).',
        func: 'x^3/40',
        weight: 0.61,
        func_mark: [
          {
            name: 'Высокая',
            func: 'x',
          },
          {
            name: 'Низкая',
            func: '1-x',
          },
          {
            name: 'Очень низкая',
            func: 'x*x',
          },
        ],
      },
      {
        definition:
          'С2 Количество конкурентов – количество похожих товаров на рынке.',
        extra_info:
          '(градаций 5: очень низкий, низкий, средний, высокий, очень высокий). Очень высокий принять равным 4.',
        func: '-(x/4)+1',
        weight: 0.15,
        func_mark: [
          {
            name: 'Высокая',
            func: 'x',
          },
          {
            name: 'Низкая',
            func: '1-x',
          },
          {
            name: 'Очень низкая',
            func: 'x*x',
          },
        ],
      },
      {
        definition:
          'С3 Аудитория – кол-во пользователей (в млн), которые используют или знают о данном бренде.',
        extra_info: '',
        func: '(-1/(x*k+1))+1',
        weight: 0.13,
        func_mark: [
          {
            name: 'Известный',
            func: 'x',
          },
          {
            name: 'Неизвестный',
            func: '1-x',
          },
        ],
      },
      {
        definition:
          'С4 Стоимость –  наиболее вероятная цена (в десятках тысяч), по которой товар или услуга могут быть проданы на свободном рынке в условиях конкуренции, когда стороны сделки действуют разумно, располагая всей необходимой информацией.',
        extra_info: '',
        func: '(x/k)^2',
        weight: 0.13,
        func_mark: [
          {
            name: 'Низкий',
            func: 'x',
          },
          {
            name: 'Очень низкий',
            func: 'x*x',
          },
          {
            name: 'Высокий',
            func: '1-x',
          },
        ],
      },
    ],
    alternative: [
      {
        description:
          'А1 Автомобиль BMW - год выпуска 2010, тип кузова Седан, количество дверей - 4, Автомат трансмиссия, мощность двигателя - 204 л.с., тип топлива - Бензин, максимальная скорость 234 км.',
        criteria_count: [
          {
            count: 'очень высокий',
            value: '',
          },
          {
            count: 'высокий',
            value: '',
          },
          {
            count: 4.36,
            value: 'млн',
          },
          {
            count: 108,
            value: 'т.руб',
          },
        ],
      },
      {
        description:
          'А2 Компьютер Sony –  11.6″, Intel Core i3, 1,3 ГГц, 4 ГБ памяти, объем SATA-диска — 500 ГБ, Bluetooth, Wi-Fi.',
        criteria_count: [
          {
            count: 'очень низкий',
            value: '',
          },
          {
            count: 'средний',
            value: '',
          },
          {
            count: 5.05,
            value: 'млн',
          },
          {
            count: 3,
            value: 'т.руб',
          },
        ],
      },
      {
        description:
          'А3 Пылесос Агрессор – цвет: синий. мощность: 100 - 2200 вт. вес без насадок: 6,1 кг. покрытие - soft touch. hepa - фильтр. объем сменного пылесборника 5,5 л.',
        criteria_count: [
          {
            count: 'низкий',
            value: '',
          },
          {
            count: 'очень низкий',
            value: '',
          },
          {
            count: 1.03,
            value: 'млн',
          },
          {
            count: 0.59,
            value: 'т.руб',
          },
        ],
      },
    ],
    rule: [
      {
        name: 'd1 : Если С1 = высокий, и С2 = высокий, и С3 = известный или С4 = очень низкий, то Y = безупречный',
      },
      {
        name: 'd2 : Если С1 = низкий, и С3 = известный, и С4 = низкий, то Y = удовлетворительный',
      },
      {
        name: 'd3 : Если С2 = очень низкий, и С3 = неизвестный или С4= высокий, то Y = неудовлетворительный',
      },
    ],
    importance_criteria: [
      {
        importance: 'важный',
        short_importance_name: 'в',
        points: [
          {
            X: 0.61,
            Y: 0,
          },
          {
            X: 0.86,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
      {
        importance: 'очень важный',
        short_importance_name: 'ов',
        points: [
          {
            X: 0.9,
            Y: 0,
          },
          {
            X: 0.11,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
      {
        importance: 'не очень важный',
        short_importance_name: 'нов',
        points: [
          {
            X: 0.23,
            Y: 0,
          },
          {
            X: 0.62,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
    ],
    importance_alternative: [
      {
        importance: 'очень низкий',
        short_importance_name: 'он',
        points: [
          {
            X: 0.12,
            Y: 0,
          },
          {
            X: 0.38,
            Y: 1,
          },
          {
            X: 0.6,
            Y: 0,
          },
        ],
      },
      {
        importance: 'низкий',
        short_importance_name: 'н',
        points: [
          {
            X: 0.32,
            Y: 0,
          },
          {
            X: 0.58,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
      {
        importance: 'средний',
        short_importance_name: 'с',
        points: [
          {
            X: 0.57,
            Y: 0,
          },
          {
            X: 0.78,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
      {
        importance: 'высокий',
        short_importance_name: 'в',
        points: [
          {
            X: 0.89,
            Y: 0,
          },
          {
            X: 0.87,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
      {
        importance: 'очень высокий',
        short_importance_name: 'ов',
        points: [
          {
            X: 0.83,
            Y: 0,
          },
          {
            X: 0.12,
            Y: 1,
          },
          {
            X: 0.9,
            Y: 0,
          },
        ],
      },
    ],
  },
};
