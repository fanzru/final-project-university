export type DetailPaper = {
  paper_id: number;
  link_pdf: string;
  len_head: number;
  body: Body[];
};

export type Body = {
  head: string;
  head_key: number;
  sentences: Sentence[];

}
export type Sentence = {
  is_important: boolean;
  sent_id: number;
  text: string;
};