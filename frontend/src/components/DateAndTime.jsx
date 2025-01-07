const formatDateTimeFrom = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = now.getMonth() + 1; // Months are zero-indexed
  const day = now.getDate();
  const hours = now.getHours();
  const minutes = now.getMinutes().toString().padStart(2, '0'); // Add leading zero if needed

  return `${year}-${month}-${day} ${hours}:${minutes}`;
};

const formatDateTimeTo = (isoString) => {
  const date = new Date(isoString);

  // Format to 'YYYY-MM-DD HH:MM' in local time
  const formatted = date.getFullYear() +
    "-" + String(date.getMonth() + 1).padStart(2, '0') +
    "-" + String(date.getDate()).padStart(2, '0') +
    " " + String(date.getHours()).padStart(2, '0') +
    ":" + String(date.getMinutes()).padStart(2, '0');

  return formatted;
}

export { formatDateTimeFrom, formatDateTimeTo };
