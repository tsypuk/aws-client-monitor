// Function to sort the table by the specified column index
function sortTable(columnIndex) {
    var table = document.getElementById("dashboardTable");
    var rows = Array.from(table.rows).slice(1); // Get all rows except the header
    var isAscending = table.getAttribute("data-sort-asc") === "true";

    rows.sort(function(rowA, rowB) {
        var cellA = rowA.cells[columnIndex].innerText;
        var cellB = rowB.cells[columnIndex].innerText;

        // Check if sorting by numbers (ID column)
        if (!isNaN(cellA) && !isNaN(cellB)) {
            return isAscending ? cellA - cellB : cellB - cellA;
        }

        // Otherwise sort by strings (e.g., Name, Status, Date)
        return isAscending ? cellA.localeCompare(cellB) : cellB.localeCompare(cellA);
    });

    // Reorder the rows in the table
    for (var i = 0; i < rows.length; i++) {
        table.tBodies[0].appendChild(rows[i]);
    }

    // Toggle the sorting order for the next click
    table.setAttribute("data-sort-asc", !isAscending);
}
