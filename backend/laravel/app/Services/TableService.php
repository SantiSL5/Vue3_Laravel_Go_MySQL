<?php

namespace App\Services;

use App\Interfaces\TableRepositoryInterface;
use App\Interfaces\CategoryRepositoryInterface;
use App\Http\Requests\Table\CreateTable;
use App\Http\Requests\Table\UpdateTable;
use App\Http\Resources\TableResource;
use App\Http\Resources\CategoryResource;

class TableService 
{
    private $tableRepository;
    private $categoryRepository;

    public function __construct(TableRepositoryInterface $tableRepository, CategoryRepositoryInterface $categoryRepository)
    {
        $this->tableRepository = $tableRepository;
        $this->categoryRepository = $categoryRepository;
    }

    public function getAllTablesService()
    {
        return TableResource::collection($this->tableRepository->getAllTablesRepo());
    }

    public function getTableByIdService($id) 
    {
        try {
            return TableResource::make($this->tableRepository->getTableByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Table doesn't exist"
            ]);
        }
    }

    public function createTableService(CreateTable $request)
    {
        try {
            CategoryResource::make($this->categoryRepository->getCategoryByIdRepo($request->get('category')));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Category doesn't exist"
            ]);
        }

        try {
            return TableResource::make($this->tableRepository->createTableRepo($request));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Table code already exists"
            ]);
        }
    }

    public function updateTableService(UpdateTable $request, $id)
    {
        try {
            TableResource::make($this->tableRepository->getTableByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Table doesn't exist"
            ]);
        }

        if ($request->get('category')) {
            try {
                CategoryResource::make($this->categoryRepository->getCategoryByIdRepo($request->get('category')));
            } catch (\Exception $e) {
                return response()->json([
                    "Message" => "Category doesn't exist"
                ]);
            }
        }

        if ($request->get('code')) {
            try {
                TableResource::make($this->tableRepository->getTableByCodeRepo($request->get('code')));
                return response()->json([
                    "Message" => "Table code already exists"
                ]);
            } catch (\Exception $e) {
                $update = $this->tableRepository->updateTableRepo($request, $id);
                if ($update == 1) {
                    return response()->json([
                        "Message" => "Updated correctly"
                    ]);
                } else {
                    return response()->json([
                        "Status" => "Not found"
                    ], 404);
                };
            }
        }
        
        try {
            $update = $this->tableRepository->updateTableRepo($request, $id);
            if ($update == 1) {
                return TableResource::make($this->tableRepository->getTableByIdRepo($id));                ;
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Table code already exists"
            ]);
        }
    }

    public function deleteTableService($id)
    {
        $delete = $this->tableRepository->deleteTableRepo($id);

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