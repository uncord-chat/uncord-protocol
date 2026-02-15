/** Server role returned by the API. */
export interface Role {
  id: string;
  name: string;
  colour: number;
  position: number;
  hoist: boolean;
  permissions: number;
  is_everyone: boolean;
  created_at: string;
  updated_at: string;
}

/** Body for POST /api/v1/server/roles. */
export interface CreateRoleRequest {
  name: string;
  colour?: number | null;
  permissions?: number | null;
  hoist?: boolean | null;
}

/** Body for PATCH /api/v1/server/roles/:roleID. */
export interface UpdateRoleRequest {
  name?: string | null;
  colour?: number | null;
  position?: number | null;
  permissions?: number | null;
  hoist?: boolean | null;
}
