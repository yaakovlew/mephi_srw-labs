export const  isMatrixNotEmpty = (matrix: any) =>  {
    // Проверяем, что матрица — это массив и содержит хотя бы одну строку
    if (Array.isArray(matrix) && matrix.length > 0) {
        // Проверяем, что хотя бы одна строка в матрице содержит элементы
        return matrix.some(row => Array.isArray(row) && row.length > 0);
    }
    return false;
}

export const isDataNotEmpty = (data:any) =>  {
    // Проверяем, что data — массив и содержит хотя бы один элемент
    if (Array.isArray(data) && data.length > 0) {
        // Проверяем, что хотя бы один элемент является объектом с ключами
        return data.some(
            item => typeof item === 'object' && item !== null && Object.keys(item).length > 0
        );
    }
    return false; // Если это не массив или массив пуст
}