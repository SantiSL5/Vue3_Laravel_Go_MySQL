import Vuex from "vuex";
import Constant from '../Constant';
import shortid from 'shortid';

export default Vuex.createStore({
    state: {
        tables: [
            { "id": shortid.generate(), "code": "A1", "category": "", "reserved": false, "capacity": 4 },
            { "id": shortid.generate(), "code": "A2", "category": "", "reserved": true, "capacity": 2 },
            { "id": shortid.generate(), "code": "D2", "category": "", "reserved": false, "capacity": 4 },
            { "id": shortid.generate(), "code": "D3", "category": "", "reserved": false, "capacity": 6 }
        ]
    },
    mutations: {
        [Constant.ADD_TABLE]: (state, payload) => {
            state.todolist.push({ ...payload.todoitem, id: shortid.generate(), done: false });
            state.todoitem = { id: "", todo: "", desc: "", done: false };
        },
        [Constant.DELETE_TABLE]: (state, payload) => {
            let index = state.todolist.findIndex((item) => item.id === payload.id);
            state.todolist.splice(index, 1);
        },
        [Constant.TOGGLE_DONE]: (state, payload) => {
            let index = state.todolist.findIndex((item) => item.id === payload.id);
            state.todolist[index].done = !state.todolist[index].done;
        },
        [Constant.UPDATE_TABLE]: (state, payload) => {
            let index = state.todolist.findIndex((item) => item.id === payload.todoitem.id);
            state.todolist[index] = payload.todoitem;
        },
        [Constant.INITIALIZE_TABLE]: (state, payload) => {
            if (payload && payload.todoitem) {
                state.todoitem = payload.todoitem;
            } else {
                state.todoitem = { id: "", todo: "", desc: "", done: false };
            }
        },
    },
    actions: {
        [Constant.ADD_TABLE]: (store, payload) => {
            store.commit(Constant.ADD_TODO, payload);
        },
        [Constant.DELETE_TABLE]: (store, payload) => {
            store.commit(Constant.DELETE_TODO, payload);
        },
        [Constant.TOGGLE_DONE]: (store, payload) => {
            store.commit(Constant.TOGGLE_DONE, payload);
        },
        [Constant.UPDATE_TABLE]: (store, payload) => {
            store.commit(Constant.UPDATE_TODO, payload);
        },
        [Constant.INITIALIZE_TODOITEM]: (store, payload) => {
            store.commit(Constant.INITIALIZE_TODOITEM, payload);
        }
    }
});