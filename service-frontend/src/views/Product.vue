<template>
    <div>
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="20" :md="18" :lg="18">
                <div>
                    <h1>Product List</h1>
                    <el-table
                            :data="productList"
                            stripe
                            border
                            @cell-click='handleCellClick'
                            style="width: 100%">

                        <el-table-column v-for="column in columns"
                                         :class-name="column.prop === 'name' ? 'column-name' : 'column-default'"
                                         :prop="column.prop" :label="column.label">
                        </el-table-column>
                    </el-table>
                </div>
            </el-col>
        </el-row>
        <el-row type="flex" justify="center" class="pagination">
            <el-col :xs="20" :sm="20" :md="18" :lg="18">
                <el-pagination
                        background
                        layout="prev, pager, next"
                        :page-size="itemCountPerPage"
                        :current-page.sync="currentPage"
                        @current-change="handleCurrentPageChange"
                        :total="totalItemCount">
                </el-pagination>
            </el-col>
        </el-row>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions, mapGetters, mapMutations, mapState } from 'vuex'
import { Exception, Product as ProductItem, Pagination } from '@/generated/swagger'
import { ProductAPI, } from '@/common/product.service.ts'

@Component({
    components: {},
    computed: {
        ...mapState([ 'uid' ]),
        ...mapGetters([ 'authenticated' ]),
    },
})
export default class Product extends Vue {
    public $refs: any
    public $notify: any
    public $router: any
    public $store: any

    public itemCountPerPage = 10
    public currentPage = 1 // offset + 1
    public totalItemCount = 4

    public columns = [
        { prop: 'name', label: 'Name', },
        { prop: 'categoryDisplayName', label: 'Category', },
        { prop: 'description', label: 'Description', },
        { prop: 'onSale', label: 'On Sale', },
        { prop: 'price', label: 'Price (KRW)', },
        { prop: 'createdAt', label: 'Registered At', },
    ]


    public productList = []

    fetchAllProducts(currentPage) {
        ProductAPI.findAll(this.itemCountPerPage, currentPage - 1, { credentials: 'include' })
            .then((response: any) => {
                const pagination: Pagination = response.pagination
                this.totalItemCount = pagination.totalItemCount
                this.productList = response.rows
            })
            .catch((response) => {
                if (!response.json) {
                    this.$notify.error({
                        title: `Error (Connection)`,
                        message: "Server is not available",
                    })

                    return
                }

                response.json().then((parsed: Exception) => {
                    this.$notify.error({
                        title: `Error (${parsed.type})`,
                        message: parsed.message,
                    })
                })
            })

    }

    mounted() {
        this.fetchAllProducts(this.currentPage)
    }

    handleCurrentPageChange(newPage) {
        this.fetchAllProducts(newPage)
    }

    handleCellClick(row, column, cell, event) {
        const productID = row.id
        const columnName = column.property

        if (columnName === 'name') {
            this.$router.push({
                name: 'product.detail',
                params: {
                    productID: productID,
                }
            })
        }
    }
}
</script>

<style>
.pagination {
    margin-top: 30px;
    text-align: center;
}

.column-name:hover {
    text-decoration: underline;
    cursor: pointer;
}

.column-default {

}
</style>