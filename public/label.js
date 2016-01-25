xxx = {}

$(document).ready(function() {
  $('#label').nutritionLabel({
    width: '300',
  });
  $("#process").submit(function(event) {
    event.preventDefault();
    $.getJSON("/process", {ingredients: $("#ingredients").val()}).done(function(data){
      var iList = $("#ingredient-list")
      iList.empty().append('<h3>Ingredients</h3>');
      for (var i in data.ingredients) {
        var ingredient = data.ingredients[i];
        switch (ingredient[0]) {
          case 0.333:
            ingredient[0] = '&frac13;';
            break;
          case 0.666:
            ingredient[0] = '&frac23;';
            break;
          case 0.125:
            ingredient[0] = '&frac18;';
            break;
          case 0.25:
            ingredient[0] = '&frac14;';
            break;
          case 0.5:
            ingredient[0] = '&frac12;';
            break;
          case 0.75:
            ingredient[0] = '&frac34;';
            break;
        }
        iList.append('<tr>' + ingredient.map(function(d) {return '<td>' + d + '</td>';}).join() + '</tr>')
      }

      $('#label').nutritionLabel({
        width: '300',
        showServingUnitQuantity : false,
        ingredientList : data.ingredients.map(function(i) {return i[2]}).join(', '),

        showPolyFat : false,
        showMonoFat : false,

        valueCalories : 450,
        valueFatCalories : 430,
        valueTotalFat : 48,
        valueSatFat : 6,
        valueTransFat : 0,
        valueCholesterol : 30,
        valueSodium : 780,
        valueTotalCarb : 3,
        valueFibers : 0,
        valueSugars : 3,
        valueProteins : 3,
        valueVitaminA : 0,
        valueVitaminC : 0,
        valueCalcium : 0,
        valueIron : 0
      });
      console.log(data);
    }).error(function(err) {
      console.log(err);
    });
    return false;
  });
});
