package Binding

type BoundITreeRewriter interface {
    RewriteStatement(BoundITreeRewriter, BoundStatement) BoundStatement 
    RewriteExpression(BoundITreeRewriter, BoundExpression) BoundExpression

    RewriteLabelStatement(BoundITreeRewriter, *BoundLabelStatement) BoundStatement
    RewriteGotoStatement(BoundITreeRewriter, *BoundGotoStatement) BoundStatement
    RewriteConditionalGotoStatement(BoundITreeRewriter, *BoundConditionalGotoStatement) BoundStatement
    RewriteBlockStatement(BoundITreeRewriter, *BoundBlockStatement) BoundStatement 
    RewriteExpressionStatement(BoundITreeRewriter, *BoundExpressionStatement) BoundStatement 
    RewriteIfStatement(BoundITreeRewriter, *BoundIfStatement) BoundStatement 
    RewriteWhileStatement(BoundITreeRewriter, *BoundWhileStatement) BoundStatement 
    RewriteForStatement(BoundITreeRewriter, *BoundForStatement) BoundStatement 
    RewriteVariableDeclaration(BoundITreeRewriter, *BoundVariableDeclaration) BoundStatement 
    RewriteUnaryExpression(BoundITreeRewriter, *BoundUnaryExpression) BoundExpression 
    RewriteLiteralExpression(BoundITreeRewriter, *BoundLiteralExpression) BoundExpression 
    RewriteBinaryExpression(BoundITreeRewriter, *BoundBinaryExpression) BoundExpression 
    RewriteVariableExpression(BoundITreeRewriter, *BoundVariableExpression) BoundExpression 
    RewriteAssignmentExpression(BoundITreeRewriter, *BoundAssignmentExpression) BoundExpression 
}
