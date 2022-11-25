<template>
    <li :class="checked(table.done)" @click="toggleDone(table.id)">
        <span :class="{ pointer:true, 'todo-done':table.reserved }" :code="'nose : ' + table.category">
            {{table.todo}}
            {{table.reserved ? " (Reserved)" : ""}}
        </span>
        <div class="float-right">
            <span class="badge badge-secondary pointer ml-1" @click.stop="editTable(table.id)">Update</span> 
            <span class="badge badge-secondary pointer ml-1" @click.stop="deleteTable(table.id)">Delete</span>
        </div>
    </li>
</template>

<script>
import Constant from '../Constant';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    props: {
        table: Object
    },
    setup(props) {
        const store = useStore();
        const router = useRouter();

        const checked = (done) => {
            return { "list-group-item":true, "list-group-item-success":done };
        }
        const toggleDone = (id) => {
            store.dispatch(Constant.TOGGLE_DONE, { id });
        }
        const deleteTable = (id) => {
            store.dispatch(Constant.DELETE_TABLE, { id });
        }
        const editTable = (id) => {
            store.dispatch(Constant.INITIALIZE_TABLE, { table: { ...props.table } });
            router.push({ name: 'updateTable', params: { id } })
        }

        return { toggleDone, deleteTable, editTable, checked }
    }
}
</script>

<style>

</style>