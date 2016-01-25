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
      var d = data.data;

      $('#label').nutritionLabel({
        width: '300',
        showServingUnitQuantity: false,
        ingredientList: data.ingredients.map(function(i) {return i[2]}).join(', '),

        showPolyFat: false,
        showMonoFat: false,

        valueCalories: d["208"],
        valueFatCalories: d["204"] * 9,
        valueTotalFat: d["204"],
        valueSatFat: d["606"],
        valueTransFat: d["605"],
        valueCholesterol: d["601"],
        valueSodium: d["307"],
        valueTotalCarb: d["205"],
        valueFibers: d["291"],
        valueSugars: d["269"],
        valueProteins: d["203"],
        valueVitaminA: d["318"],
        valueVitaminC: d["401"],
        valueCalcium: d["301"],
        valueIron: d["303"]
      });
      console.log(data);
    }).error(function(err) {
      console.log(err);
    });
    return false;
  });
});
