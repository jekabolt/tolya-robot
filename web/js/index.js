$(function() {
    $(".multiple-select").each(function() {
        $(this).hide();

        var optionsArray = [],
            selectClasses = $(this).attr("class");

        $(this)
            .find("option")
            .each(function() {
                var optionObject = {};
                optionObject.text = $(this).text();
                optionObject.value = $(this).val();
                optionObject.class = $(this).attr("class");

                optionsArray.push(optionObject);
            });

        var multipleSelectHtml =
            "<div class='multiple-select-container " +
            selectClasses +
            "'>" +
            "<ul class='multiple-select-choices'><li class='input'><input type='text' placeholder='Please select'></li></ul>" +
            "<div class='multiple-select-dropdown'><ul>";

        for (var i = 0; i < optionsArray.length; i++) {
            multipleSelectHtml +=
                "<li class='" +
                optionsArray[i].class +
                "' data-value='" +
                optionsArray[i].value +
                "'>" +
                optionsArray[i].text +
                "</li>";
        }

        multipleSelectHtml += "</ul></div></div>";

        $(multipleSelectHtml).insertAfter($(this));
    });

    // Show dropdown when input is focused
    $(".multiple-select-container").on("click", function() {
        $(this)
            .find(".multiple-select-dropdown")
            .show();
    });

    // Close dropdown when multi-select is out of focus
    $(document).on("click", function(e) {
        var $tgt = $(e.target);

        if (!$tgt.is(".multiple-select-dropdown li") &&
            !$tgt.is(".multiple-select-container input") &&
            !$tgt.is(".multiple-select-container") &&
            !$tgt.is(".multiple-select-choices")
        ) {
            $(".multiple-select-dropdown").hide();
        }
    });

    // Add item
    $(document).on(
        "click",
        ".multiple-select-dropdown li:not(.option-disabled)",
        function(e) {
            e.preventDefault();

            var optionText = $(this).text(),
                optionValue = $(this).data("value"),
                optionClass = $(this).attr("class");

            $(this).addClass("option-disabled");

            // Add item to input
            var thisInput = $(this)
                .closest(".multiple-select-container")
                .find(".input");
            $(
                "<li data-value='" +
                optionValue +
                "'>" +
                optionText +
                " <a href='#' class='remove-item'>&times;</a></li>"
            ).insertBefore(thisInput);

            // Get current value of select field
            var currentValues = $(this)
                .closest(".multiple-select-container")
                .siblings("select")
                .val();

            // Add values to select field
            currentValues.push(optionValue);

            // Assign updated value to select field
            $(this)
                .closest(".multiple-select-container")
                .siblings("select")
                .val(currentValues);

            // Make choices active, remove input placeholder
            $(".multiple-select-choices li").addClass("choice-active");
            $(".multiple-select-choices li input")
                .attr("placeholder", "")
                .focus();

            // Hide dropdown
            $(this)
                .closest(".multiple-select-dropdown")
                .hide();

            $(document).trigger("multipleSelectItemAdded", [
                optionText,
                optionValue,
                optionClass
            ]);
        }
    );

    // Remove item
    $(document).on("click", ".multiple-select-choices li a", function() {
        var optionText = $(this)
            .closest("li")
            .text(),
            optionValue = $(this)
            .closest("li")
            .data("value"),
            optionClass = $(this)
            .closest("li")
            .attr("class");

        // Re-enable option in dropdown
        $(
            ".multiple-select-dropdown li[data-value='" + optionValue + "']"
        ).removeClass("option-disabled");

        // Get current value of the select field
        var currentValues = $(this)
            .closest(".multiple-select-container")
            .siblings("select")
            .val();

        // If the removed item is in the select value, remove it
        var index = currentValues.indexOf(optionValue);

        if (index > -1) {
            currentValues.splice(index, 1);
        }

        // Re-assign the new value to the select field
        $(this)
            .closest(".multiple-select-container")
            .siblings("select")
            .val(currentValues);

        // Return placeholder to input if no choices have been made
        if (currentValues.length == 0) {
            $(this)
                .closest(".multiple-select-choices")
                .find(".input")
                .removeClass("choice-active")
                .find("input")
                .attr("placeholder", "Please select");
        }

        // Remove item from choices
        $(this)
            .closest("li")
            .remove();

        $(document).trigger("multipleSelectItemRemoved", [
            optionText,
            optionValue,
            optionClass
        ]);
    });

    //Detect item being added
    $(document).on("multipleSelectItemAdded", function(
        e,
        optionText,
        optionValue,
        optionClass
    ) {
        console.log("item has been added");
    });

    //Detect item being removed
    $(document).on("multipleSelectItemRemoved", function(
        e,
        optionText,
        optionValue,
        optionClass
    ) {
        console.log("item has been removed");
    });
});