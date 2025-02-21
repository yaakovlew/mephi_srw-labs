import {createSlice} from "@reduxjs/toolkit";

const lvl2MatrixLocalStorage = "lvl2MatrixLocalStorage"
const lvl2DataLocalStorage = "lvl2DataLocalStorage"
const lvl2ResponseResultLocalStorage = "lvl2ResponseResultLocalStorage"

export interface lvl2SliceInterface {
    matrix: string[][][];
    data: string[][];
    responseResult: string[][];
}

const initialState: lvl2SliceInterface = {
    matrix: JSON.parse(localStorage.getItem(lvl2MatrixLocalStorage) as string) || [],
    data: JSON.parse(localStorage.getItem(lvl2DataLocalStorage) as string) || [],
    responseResult: JSON.parse(localStorage.getItem(lvl2ResponseResultLocalStorage) as string) || [],
}


const lvl2Slice = createSlice({
    name: "lvl2",
    initialState,
    reducers: {
        setMatrix: (state, action) =>  {
            const {payload, index} = action.payload;
            state.matrix[index] = payload;
            localStorage.setItem(lvl2MatrixLocalStorage, JSON.stringify(state.matrix));
        },
        setData: (state, action) => {
            const {payload, index} = action.payload;
            state.data[index] = payload;
            localStorage.setItem(lvl2DataLocalStorage, JSON.stringify(state.data));
        },
        setResponseResult: (state, action) => {
            const {payload, index} = action.payload;
            state.responseResult[index] = payload;
            localStorage.setItem(lvl2ResponseResultLocalStorage, JSON.stringify(state.responseResult));
        },
    }
})


export const getMatrixLvl2 = (index) => (state: lvl2SliceInterface) => state.lvl2.matrix[index] || null;
export const getDataLvl2 = (index) => (state: lvl2SliceInterface) => state.lvl2.data[index] || null;
export const getResponseResultLvl2 = (index) => (state: lvl2SliceInterface) => state.lvl2.responseResult[index] || null;

export const {
    setMatrix,
    setData,
    setResponseResult,
} = lvl2Slice.actions;

export default lvl2Slice.reducer;