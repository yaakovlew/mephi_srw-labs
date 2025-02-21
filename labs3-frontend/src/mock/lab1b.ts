import { Lab1b } from 'src/models/lab/lab1b';

export const importanceValues: Lab1b.ImpoortanceValues[] = [
  { value: [1], definition: 'Равная важность элементов' },
  {
    value: [3],
    definition: 'Умеренное превосходство одного элемента над другим',
  },
  {
    value: [5],
    definition:
      'Существенное или сильное превосходство одного элемента над другим',
  },
  {
    value: [7],
    definition: 'Значительное превосходство одного элемента над другим',
  },
  {
    value: [9],
    definition: 'Очень сильное превосходство одного элемента над другим',
  },
  {
    value: [2, 4, 6, 8],
    definition:
      'Промежуточные решения между двумя соседними суждениями, применяются в компромиссном случае',
  },
  {
    value: [1 / 3, 1 / 5, 1 / 7, 1 / 9, 1 / 2, 1 / 4, 1 / 6, 1 / 8],
    definition:
      'Обратные величины, полученные при сравнении второго элемента с первым, означают ту или иную степень превосходства второго элемента над первым',
  },
];

export const lab1bVariants: Lab1b.Varinats = {
  1: {
    theme: 'Покупка сноуборда',
    lpr: 'Покупатель – это девушка 25 лет с зарплатой 75тыс.руб. в месяц, которая увлекается сноубордом несколько лет и имеет возмоожность пару раз в год выехать в горы.',
    criterias: {
      Экономический: [
        { title: 'Стоимость покупки', type: 'quantitative', moreBetter: false },
        { title: 'Стоимость содержания', type: 'quantitative', moreBetter: false },
        { title: 'Акции', type: 'qualitative' },
      ],
      Имиджевый: [
        { title: 'Бренд', type: 'qualitative' },
        { title: 'Внешний вид', type: 'qualitative' },
        { title: 'Узнаваемость', type: 'qualitative' },
      ],
      Технический: [
        { title: 'Жесткость сноуборда', type: 'qualitative' },
        { title: 'Форма сноуборда', type: 'qualitative' },
        { title: 'Прогиб сноуборда', type: 'qualitative' },
        { title: 'Конструкция', type: 'qualitative' },
        { title: 'Ростовка', type: 'quantitative', moreBetter: true },
        { title: 'Ширина сноуборда', type: 'quantitative', moreBetter: false },
      ],
    },
    alternatives: [
      'Jones Snowboards Hovercraft',
      'BURTON Stylus',
      'GNU Carbon Credi',
      'Flow Merc',
    ],
  },
};
