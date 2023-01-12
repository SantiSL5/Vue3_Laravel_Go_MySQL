import Constant from "../../../Constant";
import CategoryService from "../../../services/admin/CategoryService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const categoryAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_CATEGORY]: (state, payload) => {
            state.categorieslist.push({ ...payload });
        },
        [Constant.GET_ALL_CATEGORIES]: (state, payload) => {
            state.categorieslist = payload;
        },
        [Constant.GET_ONE_CATEGORY]: (state) => {
            state.allUsers;
        },
        [Constant.UPDATE_ONE_CATEGORY]: (state, payload) => {
            let index = state.categorieslist.findIndex(
                item => item.id == payload.id
            );
            let res = Object.assign({ id: payload.id }, payload.category);
            state.categorieslist[index] = res;
        },
        [Constant.DELETE_ONE_CATEGORY]: (state, payload) => {
            let index = state.categorieslist.findIndex(
                (item) => item.id === payload
            );
            state.categorieslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_CATEGORY]: (store, payload) => {
            CategoryService.createCategory(payload.category).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.status == 201) {
                    toaster.success(`Category created successfully`);
                    store.commit(Constant.CREATE_ONE_CATEGORY, data.data)
                } else {
                    toaster.error(`Category already exists`);
                }
            });
        },
        [Constant.UPDATE_ONE_CATEGORY]: (store, payload) => {
            CategoryService.updateCategory(payload.category, payload.id).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.data != "Category already exists") {
                    toaster.success(`Category updated successfully`);
                    store.commit(Constant.UPDATE_ONE_CATEGORY, payload)
                } else {
                    toaster.error(data.data);
                }
            });
        },
        [Constant.GET_ALL_CATEGORIES]: (store) => {
            CategoryService.getAllCategories().then(data => {
                store.commit(Constant.GET_ALL_CATEGORIES, data.data);
            });
        },
        [Constant.GET_ONE_CATEGORY]: (store) => {
            CategoryService.getCategoryById().then(data => {
                store.commit(Constant.GET_ONE_CATEGORY, data.data);
            });
        },
        [Constant.DELETE_ONE_CATEGORY]: (store, payload) => {
            CategoryService.deleteCategoryById(payload.id).then(
                store.commit(Constant.DELETE_ONE_CATEGORY, payload.id)
            );
        },
    },
    getters: {
        getCategory(state) {
            return state.categorieslist;
        }
    }
}
