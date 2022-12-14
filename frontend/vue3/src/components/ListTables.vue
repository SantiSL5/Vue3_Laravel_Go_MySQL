<template>
    <div class="col-3">
        <!-- <img :src="require(`@/assets/${tableitem.photo}`)" alt="Es definitivamente una fotografía" class="img-fluid"> -->
        {{ tableitem.capacity }}
        {{ tableitem.code }}
        <span v-bind:class="[tableitem.status === 'disable' ? 'badge bg-danger p-2 m-1' : 'badge bg-success p-2 m-1']"
            @click.stop="viewCategory(tableitem.id)"> {{ tableitem.status }} </span>
    </div>
</template>

<script>
export default {
    props: {
        tableitem: Object,
    }
};
</script>

<style>

</style>
