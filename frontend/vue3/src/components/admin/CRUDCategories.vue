<template>
    <h2 class="mt-2">Categories</h2>
    <form>
        <div class="mb-3">
            <label for="name" class="form-label">Name</label>
            <input type="text" class="form-control" id="name">
        </div>
        <div class="mb-3">
            <label for="photo" class="form-label">Photo</label>
            <input type="text" class="form-control" id="photo">
        </div>
        <button type="button" class="btn btn-success mb-3">Create</button>
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
                    <button type="button" @click.stop="updateCategory(category.id)"
                        class="btn btn-warning">Update</button>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'

export default {
    props: {
        categories: Array,
    },
    setup() {
        const store = useStore();

        const deleteCategory = (id) => {
            store.dispatch("categoryAdmin/" + Constant.DELETE_ONE_CATEGORY, { id });
        }
        const updateCategory = (id) => {
            store.dispatch("Category/" + Constant.INITIALIZE_Category, { id });
        }

        return { deleteCategory, updateCategory }
    },
};

</script>