import {createApi, fetchBaseQuery, RootState} from '@reduxjs/toolkit/query/react';

// const baseUrl = process.env.VITE_API_URL || '{{StageLabURL}}';

const base1BUrl = process.env.BASE_1B_URL || "https://mephi71.ru:8001";


export const lab1BApi = createApi({
    reducerPath: 'lab1BApi',
    baseQuery: fetchBaseQuery({
        baseUrl: base1BUrl,
        prepareHeaders: (headers, { getState }) => {
            const token = (getState() as RootState).auth.token;
            if (token) {
                headers.set('Authorization', `Bearer ${token}`);
            }
            
            headers.set('lab-token', 'lksdfmmskdkfmlsmdfklsdmfksdfmsldfkmsldkflsdkl');
            return headers;
        },
    }),
    endpoints: (builder) => ({
        getVariant: builder.query({
            query: ({ laboratoryId, minutesDuration }) => {
                return {
                    method: "GET",
                    url: "lab1b/variant",
                    params: {
                        laboratory_id: laboratoryId,
                        minutes_duration: minutesDuration,
                    }
                }
            },
        }),
        addCriteria: builder.mutation({
            query: ({criterias}) => ({
                url: 'lab1b/variant/criteria',
                method: 'POST',
                body: {
                    criterias: criterias
                },
            }),
        }),
        addAlternatives: builder.mutation({
            query: ({alternatives}) => ({
                url: 'lab1b/variant/alternative',
                method: 'POST',
                body: {
                    alternatives: alternatives
                },
            }),
        }),
        checkStep: builder.mutation({
            query: ({body, stepId, params}) => ({
                url: `lab1b/variant/${stepId}`,
                method: 'POST',
                body: body,
                params: params,
            }),
        }),
        checkMatrix: builder.mutation({
            query: ({body}) => ({
                url: `lab1b/variant/check_matrix`,
                method: 'POST',
                body: body,
            }),
        }),
        openNextStep: builder.mutation({
            query: (step) => ({
                url: `lab1b/variant/${step}`,
                method: 'PUT',
            })
        }),
        sendResult: builder.mutation({
            query: (body) => ({
                url: 'lab1b/variant/result',
                method: 'POST',
                body,
            }),
        }),
    }),
});

export const {
    useGetVariantQuery,
    useAddCriteriaMutation,
    useAddAlternativesMutation,
    useCheckStepMutation,
    useCheckMatrixMutation,
    useOpenNextStepMutation,
    useSendResultMutation,
} = lab1BApi;
