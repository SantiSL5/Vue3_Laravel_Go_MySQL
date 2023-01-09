import Constant from "../../../Constant";
import CategoryService from "../../../services/admin/CategoryService";

export const categoryAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_CATEGORIES]: (state, payload) => {
            state.categorieslist = payload;
        },
        [Constant.GET_ONE_CATEGORY]: (state) => {
            state.allUsers;
        },
        [Constant.DELETE_ONE_CATEGORY]: (state, payload) => {
            let index = state.categorieslist.findIndex(
                (item) => item.id === payload
            );
            state.categorieslist.splice(index, 1);
        },
    },
    actions: {
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
