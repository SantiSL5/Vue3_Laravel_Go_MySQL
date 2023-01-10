<?php

namespace App\Services;

use App\Interfaces\CategoryRepositoryInterface;
use App\Http\Requests\Category\CreateCategory;
use App\Http\Requests\Category\UpdateCategory;
use App\Http\Resources\CategoryResource;

class CategoryService 
{
    private $categoryRepository;

    public function __construct(CategoryRepositoryInterface $categoryRepository)
    {
        $this->categoryRepository = $categoryRepository;
    }

    public function getAllCategoriesService()
    {
        return CategoryResource::collection($this->categoryRepository->getAllCategoriesRepo());
    }

    public function getCategoryByIdService($id) 
    {
        try {
            return CategoryResource::make($this->categoryRepository->getCategoryByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Category doesn't exist"
            ]);
        }
    }

    public function createCategoryService(CreateCategory $request)
    {
        try {
            return CategoryResource::make($this->categoryRepository->createCategoryRepo($request));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Category already exist"
            ]);
        }
    }

    public function updateCategoryService(UpdateCategory $request, $id)
    {
        try {
            CategoryResource::make($this->categoryRepository->getCategoryByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Category doesn't exist"
            ]);
        }

        try {
            $update = $this->categoryRepository->updateCategoryRepo($request, $id);
            if ($update == 1) {
                return CategoryResource::make($this->categoryRepository->getCategoryByIdRepo($id));
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json("Category already exists");
        }
    }

    public function deleteCategoryService($id)
    {
        $delete = $this->categoryRepository->deleteCategoryRepo($id);

        if ($delete == 1) {
            return response()->json([
                "Message" => "Deleted correctly"
            ], 200);
        } else {
            return response()->json([
                "Status" => "Not found"
            ], 404);
        } 
    }
}