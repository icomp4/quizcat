<template>
    <fieldset class="filterBox">
        <legend>
            Tags
        </legend>
        <ul class="unList">
            <li class="item">
                <input type="radio" id="category1" name="category" value="category1" class="hidden-radio">
                <label for="category1">Pop Culture</label>
            </li>
            <li class="item">
                <input type="radio" id="category2" name="category" value="category2" class="hidden-radio">
                <label for="category2">Video Games</label>
            </li>
            <li class="item">
                <input type="radio" id="category3" name="category" value="category3" class="hidden-radio">
                <label for="category3">Music</label>
            </li>
            <li class="item">
                <input type="radio" id="category4" name="category" value="category4" class="hidden-radio">
                <label for="category4">TV & Movies</label>
            </li>
            
        </ul>
    </fieldset>
</template>


<style>
.filterBox {
    position: absolute;
    margin-left: 10px;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 6%;
    padding: 10px;
    border: 1px solid rgb(36, 36, 36);
    border-radius: 10px;
}
legend {
    font-size: 1rem;
    font-weight: bold;
}
.unList {
    list-style: none;
    padding: 0;
}
.item {
    font-size: small;
    margin-bottom: 10px;
}

.item :hover {
    color: rgb(43, 117, 179);
    cursor: pointer;
}
.hidden-radio {
    position: absolute;
    opacity: 0;
    pointer-events: none;
}
</style>
<script>
export default {
  data() {
    return {
    isOpen: false
    }
  }
}
</script>