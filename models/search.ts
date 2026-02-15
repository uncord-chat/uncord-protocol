/** Single search result hit returned by the API. */
export interface SearchMessageHit {
  id: string;
  channel_id: string;
  author_id: string;
  content: string;
  created_at: number;
  highlights?: string[];
}

/** Top-level search response returned inside the data envelope. */
export interface SearchResponse {
  total_count: number;
  page: number;
  per_page: number;
  hits: SearchMessageHit[];
}
