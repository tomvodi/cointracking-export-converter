package blockpit_test

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/tomvodi/cointracking-export-converter/internal/blockpit"
	"github.com/tomvodi/cointracking-export-converter/internal/common"
	"github.com/tomvodi/cointracking-export-converter/internal/interfaces/mocks"
	"github.com/tomvodi/cointracking-export-converter/internal/test"
	"testing"
)

func Test_xmlWriter_WriteTransactionsToXmlFile(t *testing.T) {
	g := NewGomegaWithT(t)
	type fields struct {
		xmlFileFactory *mocks.XMLFileFactory
		xmlFile        *mocks.XMLFile
		txConverter    *mocks.BlockpitTxConverter
		filePath       string
		transactions   []*common.CointrackingTx
		wantErr        bool
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "setting a sheet XMLHeaderLabels fails",
			prepare: func(f *fields) {
				f.xmlFileFactory.EXPECT().NewXMLFile().Return(f.xmlFile)
				f.xmlFile.EXPECT().SetSheetHeader(1, blockpit.XMLHeaderLabels).
					Return(fmt.Errorf("failed setting sheet XMLHeaderLabels"))
				f.xmlFile.EXPECT().Close().Return(nil)

				f.wantErr = true
			},
		},
		{
			name: "failed converting cointracking tx to blockpit",
			prepare: func(f *fields) {
				f.xmlFileFactory.EXPECT().NewXMLFile().Return(f.xmlFile)
				f.xmlFile.EXPECT().Close().Return(nil)

				testCtTx := test.RandomCtTx()
				f.transactions = []*common.CointrackingTx{testCtTx}

				f.txConverter.EXPECT().FromCointrackingTx(testCtTx).
					Return(nil, fmt.Errorf("failed converting"))

				f.wantErr = true
			},
		},
		{
			name: "adding sheet row fails",
			prepare: func(f *fields) {
				f.xmlFileFactory.EXPECT().NewXMLFile().Return(f.xmlFile)
				f.xmlFile.EXPECT().SetSheetHeader(1, blockpit.XMLHeaderLabels).Return(nil)
				f.xmlFile.EXPECT().Close().Return(nil)

				testCtTx := test.RandomCtTx()
				f.transactions = []*common.CointrackingTx{testCtTx}

				bpTx := test.BpTxForCtTx(testCtTx)
				f.txConverter.EXPECT().FromCointrackingTx(testCtTx).
					Return(bpTx, nil)
				f.xmlFile.EXPECT().SetSheetRow(1, 2, mock.Anything).
					Return(fmt.Errorf("failed setting sheet row"))

				f.wantErr = true
			},
		},
		{
			name: "saving sheet fails",
			prepare: func(f *fields) {
				f.xmlFileFactory.EXPECT().NewXMLFile().Return(f.xmlFile)
				f.xmlFile.EXPECT().SetSheetHeader(1, blockpit.XMLHeaderLabels).Return(nil)
				f.xmlFile.EXPECT().Close().Return(nil)

				testCtTx := test.RandomCtTx()
				f.transactions = []*common.CointrackingTx{testCtTx}

				bpTx := test.BpTxForCtTx(testCtTx)
				f.txConverter.EXPECT().FromCointrackingTx(testCtTx).
					Return(bpTx, nil)
				f.xmlFile.EXPECT().SetSheetRow(1, 2, mock.Anything).
					Return(nil)

				f.filePath = "/tmp/myfile.test"
				f.xmlFile.EXPECT().SaveAs(f.filePath).
					Return(fmt.Errorf("failed saving sheet"))

				f.wantErr = true
			},
		},
		{
			name: "saving export file succeeds",
			prepare: func(f *fields) {
				f.xmlFileFactory.EXPECT().NewXMLFile().Return(f.xmlFile)
				f.xmlFile.EXPECT().SetSheetHeader(1, blockpit.XMLHeaderLabels).Return(nil)
				f.xmlFile.EXPECT().Close().Return(nil)

				testCtTx := test.RandomCtTx()
				f.transactions = []*common.CointrackingTx{testCtTx}

				bpTx := test.BpTxForCtTx(testCtTx)
				f.txConverter.EXPECT().FromCointrackingTx(testCtTx).
					Return(bpTx, nil)
				f.xmlFile.EXPECT().SetSheetRow(1, 2, mock.Anything).
					Return(nil)

				f.filePath = "/tmp/myfile.test"
				f.xmlFile.EXPECT().SaveAs(f.filePath).
					Return(nil)

				f.wantErr = false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				xmlFileFactory: mocks.NewXMLFileFactory(t),
				xmlFile:        mocks.NewXMLFile(t),
				txConverter:    mocks.NewBlockpitTxConverter(t),
			}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			x := blockpit.NewTxXMLFileWriter(
				f.xmlFileFactory,
				f.txConverter,
			)

			err := x.WriteTransactionsToFile(f.filePath, f.transactions)

			if f.wantErr {
				g.Expect(err).Should(HaveOccurred())
			} else {
				g.Expect(err).ShouldNot(HaveOccurred())
			}
		})
	}
}
