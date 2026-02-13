/** Channel category returned by the API. */
export interface Category {
  id: string;
  name: string;
  position: number;
  created_at: string;
  updated_at: string;
}

/** Body for POST /api/v1/server/categories. */
export interface CreateCategoryRequest {
  name: string;
}

/** Body for PATCH /api/v1/categories/:categoryID. */
export interface UpdateCategoryRequest {
  name?: string | null;
  position?: number | null;
}
