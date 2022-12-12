<?php

namespace App\Repositories;
use App\Models\Category;
use App\Interfaces\CategoryRepositoryInterface;
use App\Http\Requests\Category\CreateCategory;
use App\Http\Requests\Category\UpdateCategory;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class CategoryRepository implements CategoryRepositoryInterface
{

    public function getAllCategoriesRepo()
    {
        return Category::all();
    }

    public function getCategoryByIdRepo($id) 
    {
        return Category::where('id', $id)->firstOrFail();
    }

    public function createCategoryRepo(CreateCategory $request)
    {
        return Category::create($request->validated());
    }

    public function updateCategoryRepo(UpdateCategory $request, $id)
    {
        return Category::where('id', $id)->update($request->validated());
    }

    public function deleteCategoryRepo($id)
    {
        return Category::where('id', $id)->delete();
    }
}
