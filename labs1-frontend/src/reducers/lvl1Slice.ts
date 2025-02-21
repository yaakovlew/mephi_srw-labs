import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import {TableData} from "../components/lab1b/pairwisecomparisonmatrices/TableData.tsx";

const lvl1MatrixLocalStorage = "lvl1MatrixLocalStorage"
const lvl1DataLocalStorage = "lvl1DataLocalStorage"
const lvl1ResponseResultLocalStorage = "lvl1ResponseResultLocalStorage"

export interface lvl1SliceInterface {
    matrix: string[][];
    data: string[];
    responseResult: string[];
}


const initialState: lvl1SliceInterface = {
    matrix: JSON.parse(localStorage.getItem(lvl1MatrixLocalStorage) as string),
    data: JSON.parse(localStorage.getItem(lvl1DataLocalStorage) as string),
    responseResult: JSON.parse(localStorage.getItem(lvl1ResponseResultLocalStorage) as string),
}

const lvl1Slice = createSlice({
    name: "lvl1",
    initialState,
    reducers: {
        setMatrix: (state, action) =>  {
            state.matrix = action.payload;
            localStorage.setItem(lvl1MatrixLocalStorage, JSON.stringify(state.matrix));
        },
        setData: (state, action) => {
            state.data = action.payload;
            localStorage.setItem(lvl1DataLocalStorage, JSON.stringify(state.data));
        },
        setResponseResult: (state, action) => {
            state.responseResult = action.payload;
            localStorage.setItem(lvl1ResponseResultLocalStorage, JSON.stringify(state.responseResult));
        },
    },
});

export const getMatrixLvl1 = (state: lvl1SliceInterface) => state.lvl1.matrix;
export const getDataLvl1 = (state: lvl1SliceInterface) => state.lvl1.data;
export const getResponseResultLvl1 = (state: lvl1SliceInterface) => state.lvl1.responseResult;

export const getDataLvl1Transformed = (state: lvl1SliceInterface) => transformDataLvl1(state.lvl1.data)

export const transformDataLvl1 = (data: TableData[]) => ({
    x: data.map((row) => parseFloat(row.X) || 0),
    w: data.map((row) => parseFloat(row.w) || 0),
    mw: data.map((row) => parseFloat(row.Mw) || 0),
    lambda_w: data.map((row) => parseFloat(row.Lw) || 0),
    lambda_max: parseFloat(data[0]?.Lmax || "0"),
    is: parseFloat(data[0]?.IS || "0"),
    os: parseFloat((data[0]?.OS || "0").replace("%", "")),

})

export const reverseTransformDataLvl1 = (data) => {
    const { x, w, mw, lambda_w, lambda_max, is, os } = data;

    return x.map((value, index) => ({
        key: index,
        X: x[index] ?? "",
        w: w[index] ?? "",
        Mw: mw[index] ?? "",
        Lw: lambda_w[index] ?? "",
        Lmax: index === 0 ? lambda_max ?? "" : undefined,
        IS: index === 0 ? is ?? "" : undefined,
        OS: index === 0 ? os ?? "": undefined,
    }));

};


export const {
    setMatrix,
    setData,
    setResponseResult,
} = lvl1Slice.actions;

export default lvl1Slice.reducer;
