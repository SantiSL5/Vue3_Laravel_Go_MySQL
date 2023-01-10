import Constant from "../../../Constant";
import DishTypeService from "../../../services/admin/DishTypeService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const dishTypeAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_DISHTYPE]: (state, payload) => {
            state.dishtypeslist.push({ ...payload });
        },
        [Constant.GET_ALL_DISHTYPES]: (state, payload) => {
            state.dishtypeslist = payload;
        },
        [Constant.GET_ONE_DISHTYPE]: (state) => {
            state.allUsers;
        },
        [Constant.UPDATE_ONE_DISHTYPE]: (state, payload) => {
            let index = state.dishtypeslist.findIndex(
                item => item.id == payload.id
            );
            let res = Object.assign({ id: payload.id }, payload.dishtype);
            state.dishtypeslist[index] = res;
        },
        [Constant.DELETE_ONE_DISHTYPE]: (state, payload) => {
            let index = state.dishtypeslist.findIndex(
                (item) => item.id === payload
            );
            state.dishtypeslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_DISHTYPE]: (store, payload) => {
            DishTypeService.createDishType(payload.dishtype).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.status == 201) {
                    toaster.success(`DishType created successfully`);
                    store.commit(Constant.CREATE_ONE_DISHTYPE, data.data)
                } else {
                    toaster.error(`DishType already exists`);
                }
            });
        },
        [Constant.UPDATE_ONE_DISHTYPE]: (store, payload) => {
            DishTypeService.updateDishType(payload.dishtype, payload.id).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.data != "DishType already exists") {
                    toaster.success(`DishType updated successfully`);
                    store.commit(Constant.UPDATE_ONE_DISHTYPE, payload)
                } else {
                    toaster.error(data.data);
                }
            });
        },
        [Constant.GET_ALL_DISHTYPES]: (store) => {
            DishTypeService.getAllDishTypes().then(data => {
                store.commit(Constant.GET_ALL_DISHTYPES, data.data);
            });
        },
        [Constant.GET_ONE_DISHTYPE]: (store) => {
            DishTypeService.getDishTypeById().then(data => {
                store.commit(Constant.GET_ONE_DISHTYPE, data.data);
            });
        },
        [Constant.DELETE_ONE_DISHTYPE]: (store, payload) => {
            DishTypeService.deleteDishTypeById(payload.id).then(
                store.commit(Constant.DELETE_ONE_DISHTYPE, payload.id)
            );
        },
    },
    getters: {
        getDishType(state) {
            return state.dishtypeslist;
        }
    }
}
