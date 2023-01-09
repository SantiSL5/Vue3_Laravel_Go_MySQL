<template>
    <h2 class="mt-2">Categories</h2>
    <form>
        <div class="mb-3">
            <label for="name" class="form-label">Name</label>
            <input type="text" class="form-control" id="name" v-model="state.category.name" required>
        </div>
        <div class="mb-3">
            <label for="photo" class="form-label">Photo</label>
            <input type="text" class="form-control" id="photo" v-model="state.category.photo" required>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createCategory()" v-if="state.formType == 'create'"
            :disabled="state.category.name == '' || state.category.photo == ''">Create</button>
        <div v-if="state.formType == 'update'">
            <button type="button" class="btn btn-warning mb-3" @click="updateCategory()"
                :disabled="state.category.name == '' || state.category.photo == ''">Update</button>
            <button type="button" class="btn btn-success mb-3 ms-3" @click="changeForm()">Back to create</button>

        </div>
    </form>


    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Photo</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="category in categories" :key="category.id">
                <th scope="row">{{ category.id }}</th>
                <td>{{ category.name }}</td>
                <td>{{ category.photo }}</td>
                <td>
                    <button type="button" @click.stop="deleteCategory(category.id)"
                        class="btn btn-danger me-3">Delete</button>
                    <button type="button" @click.stop="changeForm(category)" class="btn btn-warning">Update</button>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'
import { reactive } from "vue";

export default {
    props: {
        categories: Array,
    },
    setup() {
        const store = useStore();
        const state = reactive({
            category: { name: "", photo: "" },
            formType: "create",
            onUpdateCategory: null,
        });

        const changeForm = (category = null) => {
            if (category) {
                state.formType = "update";
                state.onUpdateCategory = category.id
                state.category.name = category.name;
                state.category.photo = category.photo;
            } else {
                state.formType = "create";
                state.onUpdateCategory = null;
                state.category.name = "";
                state.category.photo = "";
            }
        }

        const createCategory = () => {
            store.dispatch("categoryAdmin/" + Constant.CREATE_ONE_CATEGORY, { category: state.category })
        }
        const deleteCategory = (id) => {
            store.dispatch("categoryAdmin/" + Constant.DELETE_ONE_CATEGORY, { id });
        }
        const updateCategory = () => {
            store.dispatch("categoryAdmin/" + Constant.UPDATE_ONE_CATEGORY, { category: state.category, id: state.onUpdateCategory });
        }

        return { state, changeForm, createCategory, deleteCategory, updateCategory }
    },
};

</script>