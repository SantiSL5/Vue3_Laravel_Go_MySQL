<?php

namespace App\Interfaces;
use App\Http\Requests\Category\CreateCategory;
use App\Http\Requests\Category\UpdateCategory;

interface CategoryRepositoryInterface
{
    function getAllCategoriesRepo();
    function getCategoryByIdRepo($id);
    function createCategoryRepo(CreateCategory $request);
    function updateCategoryRepo(UpdateCategory $request,$id);
    function deleteCategoryRepo($id);
}
