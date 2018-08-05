<template>
  <div>
    <el-row type="flex" justify="center" class="category-container">
      <el-col :xs="20" :sm="20" :md="20" :lg="20">
         <span style="margin-left: 8px; margin-right: 8px;">
           <router-link :to="{name: 'product'}">
             <el-button icon="el-icon-menu" :to="'/product'">
               Products
             </el-button>
           </router-link>
         </span>
        <template v-for="(category, index) in currentCategories">
           <span style="margin-left: 8px; margin-right: 8px;">
             <i class="el-icon-arrow-right"></i>
           </span>
          <el-select v-model="currentCategories[index]"
                     placeholder="Select">
            <el-option
                    :key="optionsPerCategory[index][0].value"
                    :label="optionsPerCategory[index][0].label"
                    :value="optionsPerCategory[index][0].value">
            </el-option>
          </el-select>
        </template>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center" >
      <el-col :xs="20" :sm="20" :md="20" :lg="20">
        <div class="title-container">
          <span class="title-name">{{ productItem.name }}</span>
          <el-rate v-model="reviewRate" class="title-review-rank"></el-rate>
          <p class="title-description">
            {{ productItem.description }}
          </p>
        </div>
      </el-col>
    </el-row>
    <el-row class="detail-container">
      <el-row type="flex" justify="center">
        <el-col :xs="20" :sm="20" :md="20" :lg="20">
          <el-row>
            <el-col :xs="20" :sm="20" :md="10" :lg="10">
              <div class="image-container">
                <el-card :body-style="{ padding: '0px' }">
                  <el-carousel trigger="click" height="350px" :autoplay="false">
                    <el-carousel-item>
                      <img src="../assets/gopher2.jpg" width="100%" height="100%"/>
                    </el-carousel-item>
                    <el-carousel-item>
                      <img src="../assets/gopher1.png" width="100%" height="100%"/>
                    </el-carousel-item>
                  </el-carousel>
                </el-card>
              </div>
            </el-col>
            <el-col :xs="20" :sm="20" :md="14" :lg="14" v-if="productItem">
              <div class="price-and-option-container">
                <div class="price-container">
                  <p class="total-price-text"> Price: {{ totalPrice }} KRW </p>
                </div>

                <div class="option-container">
                  <template>
                    <el-transfer
                            @change="handleProductOptionChange"
                            :titles="['Options', 'Selected']"
                            v-model="selectedOptions"
                            :data="availableOptions">
                    </el-transfer>
                  </template>
                </div>

                <div class="button-container">
                  <el-input-number class="input-quantity" @change="handleQuantityChange"
                                   :min="1" :max="99" v-model="quantity" size="meidum" controls-position="right"></el-input-number>
                  <el-button class="button-add-to-cart">Add to Cart</el-button>
                  <el-button type="primary" class="button-buy-now">Buy now</el-button>
                </div>
              </div>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapActions, mapGetters, mapMutations, mapState } from 'vuex'
import { Exception, Product as ProductItem, ProductOption,} from '@/generated/swagger'
import { ProductAPI, } from '@/common/product.service.ts'

@Component({
    components: {},
    computed: {
        ...mapState([ 'uid' ]),
        ...mapGetters([ 'authenticated' ]),
    },
})
export default class ProductDetail extends Vue {
  public $refs: any
  public $notify: any
  public $router: any
  public $route: any
  public $store: any

  public productID: string = ""
  public productItem: ProductItem = {}
  public productOptions: Array<ProductOption> = []

  public currentCategories = []
  public optionsPerCategory = []

  public quantity: number = 1 // TODO
  public reviewRate: number = 5 // TODO


  public totalPrice: string = "0"

  public selectedOptions = []
  public availableOptions = []

  mounted() {
    this.productID = this.$route.params.productID
    ProductAPI.findOneWithOptions(this.productID, { credentials: 'include' })
        .then((response: any) => {
          this.productItem = response.product
          this.productOptions = response.options

          this.totalPrice = response.product.price
          this.availableOptions = response.options.map(option=> {
            return {
              key: option.id,
              label: `${option.name} (${option.price})`,
              disabled: false,
            }
          })

          const filtered = response.product.categoryPath.split("/").filter(x => x.trim() !== "")
          this.currentCategories = filtered

          // TODO: List all available categories for navigation
          this.optionsPerCategory = filtered.map(x => [{ label: x, value: x }])

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

  calculatePrice(selectedOptionIDList) {
      let pricePerItem = Number(this.productItem.price)
      selectedOptionIDList.map(optionID => {
          optionID = Number(optionID) - 1
          const optionPrice = Number(this.productOptions[optionID].price)
          pricePerItem += optionPrice
      })
      this.totalPrice = `${pricePerItem * this.quantity}`
  }

  handleProductOptionChange(selectedOptionIDList) {
      this.calculatePrice(selectedOptionIDList)
  }

  handleQuantityChange(quantity) {
      this.calculatePrice(this.selectedOptions)
  }
}

</script>

<style>
.detail-container {
    margin-top: 15px;
}

.image-container {
    padding-left: 15px;
    padding-right: 15px;
}

.category-container {
  margin-top: 20px;
}

.title-container {
  margin-top: 40px;
  margin-bottom: 35px;
  padding-left: 15px;
  padding-right: 15px;
}

.title-name {
  font-weight: 700;
  font-size: 30px;
}

.title-review-rank{
  display: inline-block;
  margin-left: 13px;
  margin-top: 4px;
  vertical-align: top;
}

.el-icon-star-on {
  font-size: 30px;
    margin-bottom: 5px;
}

.el-icon-star-off {
  font-size: 30px;
  margin-bottom: 5px;
}

.title-description {
  font-size: 20px;
  margin-top: 10px;
  margin-left: 10px;
  margin-bottom: 10px;
}

.price-and-option-container {
  margin-left: 40px;
}

.total-price-text {
  margin-top: 0px;
  margin-bottom: 25px;
  font-size: 25px;
}

.option-container {
  margin-top: 30px;
}

.button-container {
  margin-top: 30px;
}

.input-quantity {
  margin-right: 15px;
}

.button-buy-now {
  width: 145px;
  margin-right: 5px;
}

.button-add-to-cart {
  width: 145px;
  margin-right: 5px;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
}
</style>