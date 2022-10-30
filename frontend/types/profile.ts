export type Profile = {
  id: number;
  name: string;
  email: string;
  photo_url: string;
  created_at: string;
  deleted_at: string;
  exp: number;
  papers_users: PapersUsers[];
};

export type PapersUsers = {
  id: number;
  user_id: number;
  paper_name: string;
  link_pdf: string;
  is_done: boolean;
  created_at: string;
  sentences_labels: any;
};