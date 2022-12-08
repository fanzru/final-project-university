import exportFromJSON from 'export-from-json';

export const exportData = (Data: Array<object>, FileName: string) => {
  const data = Data;
  const fileName = FileName;
  const exportType = 'csv';
  exportFromJSON({ data, fileName, exportType });
};
