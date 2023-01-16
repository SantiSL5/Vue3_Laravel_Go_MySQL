<template>
    <nav class="nav_pag" aria-label="">
        <ul class="pagination justify-content-center">
            <li class="page-item">
                <a class="page-link" @click="changePage(state.page - 1)" aria-label="Previous">
                    <span aria-hidden="true">&laquo;</span>
                    <span class="sr-only">Previous</span>
                </a>
            </li>
            <li class="page-item" v-for="(row, id) in state.totalpages" :key="id" :class="isActive(id)"
                @click="changePage(id)">
                <a class="page-link num_pag" href="#">{{ row }}</a>
            </li>
            <li class="page-item">
                <a class="page-link" @click="changePage(state.page + 1)" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                    <span class="sr-only">Next</span>
                </a>
            </li>
        </ul>
    </nav>
</template>
<script>
import { reactive, getCurrentInstance } from 'vue';
export default {
    props: {
        totalpages: Number
    },
    setup(props) {
        const { emit } = getCurrentInstance();
        const state = reactive({
            totalpages: props.totalpages,
            page: 0
        })
        const isActive = (id) => {
            return { active: id == state.page };
        };
        const changePage = (page) => {
            if (page < 0) {
                state.page = 0;
            } else if (page == state.totalpages) {
                state.page = state.totalpages - 1;
            } else {
                state.page = page;
            }
            emit('changepage', state.page)
        }
        return { state, changePage, isActive }
    }
}
</script>
