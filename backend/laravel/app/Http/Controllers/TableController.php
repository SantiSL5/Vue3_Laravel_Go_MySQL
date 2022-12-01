<?php

namespace App\Http\Controllers;

use App\Http\Requests\Table\CreateTable;
use App\Http\Requests\Table\UpdateTable;
use App\Http\Resources\TableResource;
use App\Models\Table;
use App\Models\Category;

use Illuminate\Http\JsonResponse;

class TableController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return TableResource::collection(Table::all());
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \App\Http\Requests\StoreTableRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateTable $request)
    {
        try {
            Category::where('id', $request->get('category'))->firstOrFail();
        } catch (\Exception $e) {
            return response()->json("Category doesn't exist");
        }

        try {
            return TableResource::make(Table::create($request->validated()));
        } catch (\Exception $e) {
            return response()->json("Table code already exists");
        }
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        try {
            return TableResource::make(Table::where('id', $id)->firstOrFail());
        } catch (\Exception $e) {
            return response()->json("Table doesn't exist");
        }
        return TableResource::make(Table::where('id', $id)->firstOrFail());
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function edit(Table $table)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\UpdateTableRequest  $request
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateTable $request, $id)
    {
        if ($request->get('category')) {
            try {
                Category::where('id', $request->get('category'))->firstOrFail();
            } catch (\Exception $e) {
                return response()->json("Category doesn't exist");
            }
        }

        if ($request->get('code')) {
            try {
                Table::where('code', $request->get('code'))->firstOrFail();
                return response()->json("Table code already exists");
            } catch (\Exception $e) {
                $update = Table::where('id', $id)->update($request->validated());
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
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        $delete = Table::where('id', $id)->delete();

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
